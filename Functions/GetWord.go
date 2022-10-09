package hangman

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

func GetWord() string {
	data, err := os.ReadFile("WordsLists/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	os := runtime.GOOS //runtime.GOOS -> linux, windows, darwin etc.
	println("You are running on", os, "!")
	fmt.Println("")

	var WordsList []string
	var LettersOfTheWord string
	for _, l := range data {
		if string(l) == "\n" {
			if os == "windows" {
				WordsList = append(WordsList, LettersOfTheWord[:len(LettersOfTheWord)-1]) // -1 getting rid of the backslash (byte '13')
			} else {
				WordsList = append(WordsList, LettersOfTheWord[:len(LettersOfTheWord)])
			}
			LettersOfTheWord = ""
		} else {
			LettersOfTheWord = LettersOfTheWord + string(l)
		}
	}
	rand.Seed(time.Now().UnixNano())
	word := strings.ToUpper(WordsList[rand.Intn(len(WordsList))])
	return word
	//return "EAU"
}
