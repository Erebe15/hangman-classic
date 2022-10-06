package hangman

import (
	"log"
	"math/rand"
	"os"
)

func GetWord() string {
	data, err := os.ReadFile("WordsLists/words.txt")
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
	index := rand.Int(0, len(WordsList))
	return WordsList[index]
}
