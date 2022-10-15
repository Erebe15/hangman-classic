package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

func GetWord() string {
	if len(os.Args) == 1 {
		fmt.Println("Please select a file as argument")
		os.Exit(1)
	}
	data, err := os.ReadFile(string(os.Args[1]))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Words := strings.Split(string(data), "\n")
	os := runtime.GOOS //runtime.GOOS -> linux, windows, darwin etc.
	println("*DEBUG* You are running on", os, "!")
	if os == "windows" {
		for i, word := range Words {
			Words[i] = word[:len(word)-1]
		}
	}
	rand.Seed(time.Now().UnixNano())
	//word := strings.ToUpper(Words[rand.Intn(len(Words))])
	return "TOURNESOL"

	//	data, err := os.ReadFile("FilesAndLists/words.txt")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	os := runtime.GOOS //runtime.GOOS -> linux, windows, darwin etc.
	//	println("You are running on", os, "!")
	//	fmt.Println("")
	//
	//	var WordsList []string
	//	var LettersOfTheWord string
	//	for _, l := range data {
	//		if string(l) == "\n" {
	//			if os == "windows" {
	//				WordsList = append(WordsList, LettersOfTheWord[:len(LettersOfTheWord)-1]) // -1 getting rid of the backslash (byte '13')
	//			} else {
	//				WordsList = append(WordsList, LettersOfTheWord[:len(LettersOfTheWord)])
	//			}
	//			LettersOfTheWord = ""
	//		} else {
	//			LettersOfTheWord = LettersOfTheWord + string(l)
	//		}
	//	}
	//	rand.Seed(time.Now().UnixNano())
	//	word := strings.ToUpper(WordsList[rand.Intn(len(WordsList))])
	//	return word
	//	//return "EAU"
}
