package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	// Code to measure
	words := readInWords("cmd/words.txt")
	shortList := removeWordsWithDuplicateLetters(words)
	fmt.Println(words[0])
	fmt.Println(shortList[0])
	fmt.Println(len(shortList))
	wordPairs := convertToBinary(shortList)
	wordPairs = randomizeSlice(wordPairs)
	for i := 0; i < 10; i++ {
		fmt.Printf("\n %v is %26.26b in binary", wordPairs[i].word, wordPairs[i].binary)
	}

	duration := time.Since(start)

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)
}

func randomizeSlice(a []wordPair) []wordPair {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func convertToBinary(words []string) []wordPair {
	wordPairs := []wordPair{}
	for _, word := range words {
		wordPairs = append(wordPairs, wordPair{word: word, binary: encode5(word)})
	}
	return wordPairs
}

func readInWords(filename string) []string {
	file, err := os.Open(filename)

	var words []string

	if err != nil {
		log.Fatal(err)
	}

	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan() {
		if len(Scanner.Text()) == 5 {
			words = append(words, Scanner.Text())
		}
	}

	if err := Scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

func duplicateLetterCheck(word []rune) bool {
	for i, first := range word {
		for j, second := range word {
			if i != j {
				if first == second {
					return true
				}
			}
		}
	}
	return false
}

func getLetters(word string) []rune {
	letters := []rune{}
	for _, v := range word {
		letters = append(letters, v)
	}
	return letters
}

func removeWordsWithDuplicateLetters(words []string) []string {
	wordList := []string{}
	for _, v := range words {
		if !duplicateLetterCheck(getLetters(v)) {
			wordList = append(wordList, v)
		}

	}
	// fmt.Println(letterFreq)
	return wordList
}

func encode5(s string) uint32 {
	return 1<<(s[0]-'a') + 1<<(s[1]-'a') + 1<<(s[2]-'a') + 1<<(s[3]-'a') + 1<<(s[4]-'a')
}

type wordPair struct {
	word   string
	binary uint32
}
