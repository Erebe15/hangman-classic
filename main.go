package main

import (
	"fmt"
	"hangman/Functions"
)

type game struct {
	Word                   string
	tries                  int
	guess, RevealedLettres []string
}

func main() {
	var GameInProgress game
	GameInProgress.Word = hangman.GetWord()
	fmt.Println("the secret word is:", GameInProgress.Word)
	GameInProgress.RevealedLettres = hangman.RevealStartLettres(GameInProgress.Word)
	fmt.Printf("Revealed lettres are: %s\n", GameInProgress.RevealedLettres)
}
