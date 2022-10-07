package hangman

import (
	"math/rand"
	"time"
)

func RevealStartLettres(WordChosen string) []string { // select random lettres from the word for the start
	rand.Seed(time.Now().UnixNano())
	var ChosenLettres []string
	NumberOfRevealed := (len(WordChosen) / 2) - 1

	if len(WordChosen) >= 3 {
		for len(ChosenLettres) < NumberOfRevealed {
			RndLettre := string(WordChosen[rand.Intn(len(WordChosen))])
			if !DoesContain(ChosenLettres, RndLettre) {
				ChosenLettres = append(ChosenLettres, RndLettre)
			}
		}
	}
	return ChosenLettres
}
