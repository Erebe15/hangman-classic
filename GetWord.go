package main

import (
	"fmt"
	"hangman/Functions"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetWord() string {
	rand.Seed(time.Now().UnixNano())
	sep := []byte{13, 10}
	data, err := os.ReadFile("FilesAndLists/WordLists/" + GameInProgress.set.Words)
	if err != nil {
		hangman.Clear()
		fmt.Println("error:", err)
		os.Exit(1)
	}
	Words := strings.Split(string(data), string(sep))
	word := strings.ToUpper(Words[rand.Intn(len(Words))])
	return word
}
