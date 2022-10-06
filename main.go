package main

import (
	"fmt"
	hangman "hangman/Functions"
)

type game struct {
	Advancment, Word string
	tried            int
}

func main() {
	var partie game
	partie.Word = hangman.GetWord()
	fmt.Println(partie.Word)
}
