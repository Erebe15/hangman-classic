package main

import (
	"fmt"
	"hangman/Functions"
)

type Game struct {
	Word                   string
	Tries                  int
	guess, RevealedLettres []string
}

func main() {
	var GameInProgress Game
	GameInProgress.Word = hangman.GetWord()
	fmt.Println("the secret word is:", GameInProgress.Word)
	GameInProgress.RevealedLettres = hangman.RevealStartLettres(GameInProgress.Word)
	fmt.Printf("Revealed lettres are: %s\n", GameInProgress.RevealedLettres)
	GameInProgress.Tries = 0
	StartPlaying(GameInProgress)
}
