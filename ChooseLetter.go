package main

import (
	"fmt"
	"hangman/Functions"
)

func ChooseLetter(GameInProgress Game) {

	fmt.Printf("Choose a letter : ")
	var try string
	fmt.Scanln(&try)
	GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, try)
	fmt.Printf("Revealed lettres are: %s\n", GameInProgress.RevealedLettres)
	hangman.PrintWord(GameInProgress.Word, GameInProgress.RevealedLettres)

	/*for i := 0; i < len(Word)-1; i++ {
		if DoesContain(starter, Word) == true {
			fmt.Print(string(Word[i]))
		} else {
			fmt.Print("_")

		}
	}
	*/
}
