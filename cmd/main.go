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
	duration := time.Since(start)
	letterToBinary(getLetters("a")[0])
	letterToBinary(getLetters("b")[0])
	letterToBinary(getLetters("c")[0])

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)
}

func convertToBinary(words []string) []uint32 {
	/*
		for _, word := range words {
			runes := getLetters(word)
		}
	*/
	return []uint32{}
}

func letterToBinary(letter rune) uint32 {
	var binary uint32 = 3202
	switch letter {

	case 97:
		fmt.Println("100000000000000000000000000000000000")
	case 98:
		fmt.Println("010000000000000000000000000000000000")
	case 99:
		fmt.Println("001000000000000000000000000000000000")
	default:
		fmt.Println("butts")
	}

	return binary
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

/*

package main

import (
	"fmt"
)

func main() {
	letter := "f"
	boo := math(letter)
	a := math("a")
	b := math("j")
	fmt.Printf("\n %b  %b   %b", a, b, a+b)
	fmt.Printf("\n the letter %v is %b in binary", letter, boo)
}

func math(letter string) uint64 {

	a_rune := []rune("a")[0]
	letter_rune := []rune(letter)[0]
	var x int = 1 << (letter_rune - a_rune)
	fmt.Println(x)

	fmt.Printf("\n  %b ", x)
	return uint64(x)
}


*/
