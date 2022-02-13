package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/pkg/errors"
)

func main() {
	acceptedWordFilename := "accepted.txt"

	potentialWords, err := readWords("input-fr.txt")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	acceptedWordsRaw, err := readWords(acceptedWordFilename)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	err = selectWords(potentialWords, acceptedWordsRaw, acceptedWordFilename, 5)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
}

func selectWords(potentialWords []string, acceptedWords []string, outputFilename string, wordLength int) error {
	acceptedWordsMap := map[string]bool{}
	for _, w := range acceptedWords {
		acceptedWordsMap[w] = true
	}

	wordsToTest := []string{}
	for _, word := range potentialWords {
		if len(word) == wordLength && !acceptedWordsMap[word] {
			wordsToTest = append(wordsToTest, word)
		}
	}

	fmt.Printf("Please say yes (y) if you recognize the word or no (n) if you do not\n")
	for true {
		wordToTest := selectWord(wordsToTest, acceptedWordsMap)
		if getBoolFromUser(wordToTest) {
			err := addAcceptedWord(wordToTest, outputFilename)
			if err != nil {
				return err
			}
			acceptedWordsMap[wordToTest] = true
		}
	}

	return nil
}

func addAcceptedWord(word string, outputFile string) error {
	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "unable to open output file")
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s\n", word))
	if err != nil {
		return errors.Wrapf(err, "unable to append to output file")
	}

	return nil
}

func selectWord(words []string, acceptedWords map[string]bool) string {
	targetWord := ""
	for targetWord == "" {
		index := int(rand.Float64() * float64(len(words)))
		targetWord = words[index]
		if acceptedWords[targetWord] {
			targetWord = ""
		}
	}

	return targetWord
}

func readWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open source file")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read lines")
	}
	return lines, nil
}

func getBoolFromUser(message string) bool {
	fmt.Println(message)
	answer := captureWord()

	return answer == "yes" || answer == "y"
}

func captureWord() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
