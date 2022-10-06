package hangman

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func GetWord() string {
	data, err := os.ReadFile("WordsLists/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	var WordsList []string
	var LettersOfTheWord string
	for _, l := range data {
		if string(l) == "\n" {
			WordsList = append(WordsList, LettersOfTheWord)
			LettersOfTheWord = ""
		} else {
			LettersOfTheWord = LettersOfTheWord + string(l)
		}
	}
	rand.Seed(time.Now().UnixNano())
	return WordsList[rand.Intn(len(WordsList))]
}
