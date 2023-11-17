package counter

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestUserKeyWordsCount(t *testing.T) {
	testCases := []struct {
		description string
		keywords    []string
		expected    string
	}{
		{
			description: "Test Case 1",
			keywords:    []string{"master", "again", "with", "you", "give"},
			expected: "master: 1365\n" +
				"again: 1365\n" +
				"with: 1365\n" +
				"you: 2730\n" +
				"give: 1365\n" +
				"всего: 8190",
		},
		{
			description: "Test Case 2",
			keywords:    []string{"from", "great", "has", "him", "how"},
			expected: "from: 1365\n" +
				"great: 2730\n" +
				"has: 2730\n" +
				"him: 1365\n" +
				"how: 2730\n" +
				"всего: 10920",
		},
		{
			description: "Test Case 3",
			keywords:    []string{"take", "that", "there", "toil", "truth"},
			expected: "take: 4095\n" +
				"that: 4095\n" +
				"there: 1365\n" +
				"toil: 1365\n" +
				"truth: 1365\n" +
				"всего: 12285",
		},
	}

	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to find current directory: %v", err)
	}

	textFileName := "input.txt"
	textFilePath := filepath.Join(currentDir, "../../data/input.txt")
	file, err := os.Open(textFilePath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v\n", textFileName, err)
	}
	defer file.Close()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			userKeyWordsCounter := NewUserKeyWordsCounter(tc.keywords)

			taskChan := make(chan string)
			var wg sync.WaitGroup
			for i := 0; i < 2; i++ {
				wg.Add(1)
				go Worker(taskChan, &wg, userKeyWordsCounter)
			}

			go func() {
				defer close(taskChan)
				err := UserKeyWordsCountFile(file, taskChan)
				if err != nil {
					log.Fatalf("Failed to process file: %v", err)
				}
			}()

			wg.Wait()
			result := UserKeyWordsCount(userKeyWordsCounter.UserKeyWordsCount, tc.keywords)

			if result != tc.expected {
				t.Errorf("Expected keywordCount %s, but got %s", tc.expected, result)
			}

			_, err = file.Seek(0, 0)
			if err != nil {
				t.Fatalf("Error while moving pointer in file: %v", err)
			}
		})
	}
}
