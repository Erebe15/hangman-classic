package hangman

import (
	"log"
	"math/rand"
	"os"
)

func GetWord() string {
	data, err := os.ReadFile("Words.txt")
	if err != nil {
		log.Fatal(err)
	}
	var WordsList []string
	var LettersOfTheWord string
	for _, l := range data {
		if l == '\n' {
			WordsList = append(WordsList, string(l))
		} else {
			LettersOfTheWord = LettersOfTheWord + string(l)
		}
	}
	index := rand.Intn(len(WordsList))
	return WordsList[index]
}
