package main

import (
	"bufio"
	"fmt"
	"github.com/NRKA/Parallel-Word-Counter/pkg/counter"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	fmt.Print("Please enter words separated by spaces: ")
	inputReader := bufio.NewReader(os.Stdin)
	userInput, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read user input %v", err)
	}
	userInput = strings.ToLower(strings.TrimSpace(userInput))
	userKeyWords := strings.Split(userInput, " ")

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to find current directory: %v", err)
	}
	textFileName := "input.txt"
	textFilePath := filepath.Join(currentDir, "../data/input.txt")

	file, err := os.Open(textFilePath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v\n", textFileName, err)
	}
	defer file.Close()

	userKeyWordsCounter := counter.NewUserKeyWordsCounter(userKeyWords)
	taskChan := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go counter.Worker(taskChan, &wg, userKeyWordsCounter)
	}

	go func() {
		defer close(taskChan)
		err := counter.UserKeyWordsCountFile(file, taskChan)
		if err != nil {
			log.Fatalf("Failed to process file: %v", err)
		}
	}()

	wg.Wait()
	result := counter.UserKeyWordsCount(userKeyWordsCounter.UserKeyWordsCount, userKeyWords)
	fmt.Println(result)
}
