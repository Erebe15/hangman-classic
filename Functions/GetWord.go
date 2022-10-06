package hangman

import (
	"log"
	"math/rand"
	"os"
	"strings"
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
			WordsList = append(WordsList, LettersOfTheWord[:len(LettersOfTheWord)-1]) // -1 getting rid of the backslash (byte '13')
			LettersOfTheWord = ""
		} else {
			LettersOfTheWord = LettersOfTheWord + string(l)
		}
	}
	rand.Seed(time.Now().UnixNano())
	word := strings.ToUpper(WordsList[rand.Intn(len(WordsList))])
	for i, l := range word {
		println(i, "= ", string(l))
	}
	return word
	//return "EAU"
}
