package main

import (
	"fmt"
	"hangman/Functions"
	"strings"
)

func ChooseLetter(GameInProgress Game) string {
	var letter string
	letterValid := false
	var choice string
	for !letterValid {
		fmt.Printf("Choose a letter: ")
		fmt.Scanln(&letter)
		letter = strings.ToUpper(letter)
		if !hangman.IsAlpha(letter) {
			fmt.Printf("Please, select a letter, try again.\n\n")
			fmt.Println("___________________________________")
		} else if len(letter) > 1 {
			fmt.Printf("do you think it's %s?\n", letter)
			fmt.Printf("yes / no\n")
			fmt.Scanln(&choice)
			if strings.ToUpper(choice) == "YES" {
				letterValid = true
			}
			fmt.Println("___________________________________")
		} else if hangman.DoesContain(GameInProgress.RevealedLettres, letter) {
			fmt.Printf("This letter has already been Revealed\n\n")
			fmt.Println("___________________________________")
		} else if hangman.DoesContain(GameInProgress.guess, letter) {
			fmt.Printf("This letter has already been chosen\n\n")
			fmt.Println("___________________________________")
		} else {
			letterValid = true
			GameInProgress.guess = append(GameInProgress.guess, letter)
		}
	}
	return strings.ToUpper(letter)
}
