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
	deAnagram := removeAnagrams(wordPairs)
	for _, pairs := range deAnagram {
	// check(deAnagram)
	data := makePairs(deAnagram)
	reducedPairs := removeDupePairs(data)
	pairedPairs := pairPairs(reducedPairs)
	reducedPairedPairs := removeDupePairs(pairedPairs)
	final := lastWord(reducedPairedPairs, deAnagram)
	fmt.Println(len(final))
	// fmt.Println(final[0])
	fmt.Println(final)
	duration := time.Since(start)

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Printf("\n %v \n", duration)
}

func lastWord(pairedPairs []words2, words []wordPair) []words2 {
	fiveLetter := []words2{}
	for _, pairs := range pairedPairs {
		for _, word := range words {
			if (pairs.letters & word.binary) == 0 {
				// fmt.Println("Added")
				temp := words2{}
				temp.num = 5
				temp.words = append(temp.words, pairs.words...)
				temp.words = append(temp.words, word)
				// fmt.Println("words", temp.words)
				temp.letters = pairs.letters + word.binary
				fiveLetter = append(fiveLetter, temp)
			}
		}
	}
	return fiveLetter
}

func removeAnagrams(words []wordPair) []wordPair {
	anagramLess := []wordPair{}
	for _, word := range words {
		if anagramCheck(anagramLess, word) {
			anagramLess = append(anagramLess, word)
		}
	}
	fmt.Println(len(words), len(anagramLess))
	return anagramLess
}

func anagramCheck(words []wordPair, checkWord wordPair) bool {
	for _, listEntry := range words {
		vowels := encode5("aeiou")
		if listEntry.binary == checkWord.binary   || (listEntry.binary & vowels) == 0 {
			return false
		}
	}
	return true
}

type words2 struct {
	words   []wordPair
	letters uint32
	num     int
}

func check(words []wordPair) {
	fiveLetter := []words2{}
	for i := range words {
		working := createSlice(words, i)
		fiveLetter = append(fiveLetter, checkSingleWord(words, working)...)
		// fmt.Println(i)
	}
	fmt.Println(fiveLetter[0].words)
	fmt.Println(len(fiveLetter))
	fmt.Println(fiveLetter[0].words)
	temp := []words2{}
	for i, v := range fiveLetter {
		if v.num == 5 {
			// fmt.Println(v.words)
			temp = append(temp, fiveLetter[i])
		}
	}
	fmt.Println(temp[len(temp)-1])
}

func removeDupePairs(pairs []words2) []words2 {
	duplicateLess := []words2{}
	used := make(map[uint32]bool)
	for _, pair := range pairs {
		if used[pair.letters] == false {
			duplicateLess = append(duplicateLess, pair)
			used[pair.letters] = true
		}
	}
	fmt.Println(len(pairs), len(duplicateLess))
	return duplicateLess

}

func dupeCheck(pairs []words2, checkWord words2) bool {
	for _, pair := range pairs {
		if (pair.letters & checkWord.letters) == 1 {
			return false
		}
	}
	return true
}

// fmt.Println(fiveLetter)

/*
Check the first word
go to the second level on the first word
if there is a match on the second level go to the third
if not go back up
*/

func checkSingleWord(words []wordPair, working []words2) []words2 {
	for i, entry := range working {
		for _, word := range words {
			if (entry.letters & word.binary) == 0 {
				working[i].words = append(working[i].words, word)
				working[i].letters = working[i].letters + word.binary
				working[i].num++
			}
		}
		fmt.Println(i)
	}
	temp := []words2{}
	for _, v := range working {
		if v.num == 5 {
			temp = append(temp, v)
		}
	}
	return temp
}

func pairPairs(pairs []words2) []words2 {
	pairedPairs := []words2{}
	for _, pairOne := range pairs {
		// fmt.Println(i)
		for _, pairTwo := range pairs {
			if (pairOne.letters & pairTwo.letters) == 0 {
				temp := words2{}
				temp.num = 4
				temp.words = append(temp.words, pairOne.words...)
				temp.words = append(temp.words, pairTwo.words...)
				temp.letters = pairOne.letters + pairTwo.letters
				pairedPairs = append(pairedPairs, temp)
			}
		}
	}
	return pairedPairs
}

func makePairs(words []wordPair) []words2 {
	pairs := []words2{}
	for _, wordOne := range words {
		for _, wordTwo := range words {
			if (wordOne.binary & wordTwo.binary) == 0 {
				temp := words2{}
				temp.words = append(temp.words, wordOne)
				temp.words = append(temp.words, wordTwo)
				temp.num = 2
				temp.letters = wordOne.binary + wordTwo.binary
				pairs = append(pairs, temp)
			}
		}
	}
	fmt.Println(len(pairs))
	return pairs
}

func createSlice(words []wordPair, num int) []words2 {
	first := words[num]
	temp := []words2{}
	for range words {
		temp = append(temp, words2{words: []wordPair{first}, letters: first.binary, num: 1})
	}
	return temp
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

func compare() {
	w1 := encode5("abcde")
	w2 := encode5("fghij")
	fmt.Println()
	if (w1 & w2) == 0 {
		fmt.Println("no letters in common")
	} else {
		fmt.Println("letters in common")
	}
}
