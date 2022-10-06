package main

import (
	"fmt"
	"hangman/Functions"
)

type game struct {
	Advancment, Word string
	tried            int
}

func main() {
	var partie game
	partie.Word = hangman.GetWord()
	fmt.Println(hangman.RevealWord(partie.Word))
}
