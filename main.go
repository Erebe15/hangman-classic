package main

import (
	"fmt"
	"hangman/Functions"
	"strings"
)

type Game struct {
	Word                               string
	Tries                              int
	guess, RevealedLettres, JoseStates []string
}

func main() {
	var GameInProgress Game
	PlayAgain := true
	choice := ""
	for PlayAgain {
		GameInProgress.Word = hangman.GetWord()
		fmt.Println("the secret word is:", GameInProgress.Word)
		GameInProgress.RevealedLettres = hangman.RevealStartLettres(GameInProgress.Word)
		fmt.Printf("Revealed lettres are: %s\n", GameInProgress.RevealedLettres)
		GameInProgress.Tries = 0
		GameInProgress.JoseStates = hangman.GetJose()
		hangman.PrintRules()
		StartPlaying(GameInProgress)

		fmt.Println("")
		fmt.Println("Do you want to play again ?")
		fmt.Printf("Enter 'y' to play again, or any other input to quit : ")
		fmt.Scanln(&choice)
		if strings.ToUpper(choice) == "Y" {
			PlayAgain = true
		} else {
			fmt.Println("See you later!")
			PlayAgain = false
		}
	}
}
