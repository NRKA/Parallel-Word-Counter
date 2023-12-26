package counter

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
	"sync"
)

type UserKeyWordsCounter struct {
	mu                sync.Mutex
	UserKeyWordsCount map[string]*int
}

func NewUserKeyWordsCounter(userKeyWords []string) *UserKeyWordsCounter {
	counter := &UserKeyWordsCounter{
		mu:                sync.Mutex{},
		UserKeyWordsCount: make(map[string]*int),
	}
	for _, keyWord := range userKeyWords {
		v := 0
		counter.UserKeyWordsCount[keyWord] = &v
	}
	return counter
}

func (counter *UserKeyWordsCounter) ProcessLine(line string) {
	lineWords := strings.Fields(line)

	for _, word := range lineWords {
		word = strings.ToLower(word)
		for keyword, value := range counter.UserKeyWordsCount {
			count := strings.Count(word, keyword)
			if count > 0 {
				counter.mu.Lock()
				*value += count
				counter.mu.Unlock()
			}

		}
	}
}

func Worker(taskQueue <-chan string, wg *sync.WaitGroup, userKeyWordsCounter *UserKeyWordsCounter) {
	defer wg.Done()
	for line := range taskQueue {
		userKeyWordsCounter.ProcessLine(line)
	}
}

func UserKeyWordsCountFile(reader io.Reader, taskChan chan<- string) error {
	fileReader := bufio.NewReader(reader)
	for {
		line, err := fileReader.ReadString('\n')
		isEOF := errors.Is(err, io.EOF)
		if err != nil && !isEOF {
			return err
		}
		taskChan <- line
		if isEOF {
			break
		}
	}
	return nil
}

func UserKeyWordsCount(wordCount map[string]*int, userKeyWords []string) string {
	var (
		result strings.Builder
		total  int
	)
	for _, str := range userKeyWords {
		result.WriteString(str + ":" + " " + strconv.Itoa(*wordCount[str]) + "\n")
		total += *wordCount[str]
	}
	result.WriteString("total: " + strconv.Itoa(total))
	return result.String()
}
