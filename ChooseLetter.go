package main

import (
	"fmt"
	hangman "hangman/Functions"
	"strings"
)

func ChooseLetter(GameInProgress Game) string {
	var letter string
	//var letterFalse []string
	letterValid := false
	for !letterValid {
		fmt.Printf("Choose a letter: ")
		fmt.Scanln(&letter)
		letter = strings.ToUpper(letter)
		if !hangman.IsAlpha(letter) {
			fmt.Println("Please, select a letter, try again. ")
			fmt.Println("")
			fmt.Println("___________________________________")
		} else if len(letter) > 1 {
			fmt.Println("Please, select only one letter, try again. ")
			fmt.Println("")
			fmt.Println("___________________________________")
		} else if hangman.DoesContain(GameInProgress.RevealedLettres, letter) {
			fmt.Println("This letter has already been Revealed")
			fmt.Println("")
			fmt.Println("___________________________________")
		} else if hangman.DoesContain(GameInProgress.guess, letter) {
			fmt.Println("This letter has already been chosen")
			fmt.Println("")
			fmt.Println("___________________________________")
		} else {
			letterValid = true
			println("letter: ", letter)
			GameInProgress.guess = append(GameInProgress.guess, letter)
		}
	}
	return letter
}
