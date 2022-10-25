package main

import (
	"fmt"
	"hangman/Functions"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func GetWord() string {
	if GameInProgress.set.Words == "слова.txt" && GameInProgress.set.Difficulty >= 4 {
		GameInProgress.set.Words = "words.txt"
	} else if GameInProgress.set.Difficulty >= 4 {
		GameInProgress.set.Words = "слова.txt"
	}
	rand.Seed(time.Now().UnixNano())
	sep := []byte{13, 10}
	data, err := os.ReadFile("FilesAndLists/WordLists/" + GameInProgress.set.Words)
	if err != nil {
		hangman.Clear()
		fmt.Println("error:", err)
		os.Exit(1)
	}
	var word string
	Words := strings.Split(string(data), string(sep))
	word = strings.ToUpper(Words[rand.Intn(len(Words))])
	for utf8.RuneCountInString(word) < 8 && GameInProgress.set.Difficulty >= 3 {
		word = strings.ToUpper(Words[rand.Intn(len(Words))])
	}
	return word
}
