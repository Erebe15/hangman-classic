package main

import (
	"fmt"
	"hangman/Functions"
	"os"
	"strings"
	"time"
)

func ChooseLetter() string {
	var guess string
	ValidLetter := false
	var choice string

	for !ValidLetter {
		fmt.Scan(&guess)
		guess = strings.ToUpper(guess)

		switch {

		case guess == "STOP":
			docName := ""
			fmt.Printf("\x1B[CPlease the the save name : ")
			fmt.Scanln(&docName)
			GameInProgress.Status += 10
			Marshal("Saves/" + docName + ".json")
			fmt.Printf("\x1B[CGame saved as : %s", docName)
			time.Sleep(time.Second * 2)
			hangman.Clear()
			os.Exit(0)

		case !hangman.IsAlpha(guess):
			fmt.Print("\a\x1B[C\x1B[31mPlease, select a letter\x1B[0m\n\x1B[C")
			break

		case len(guess) > 1:
			ClearTerminal()
			fmt.Printf("do you think it's %s?\n\x1B[C\x1B[32myes\x1B[0m / \x1B[31mno\x1B[0m\n\x1B[C", guess)
			fmt.Scanln(&choice)
			if strings.ToUpper(choice) == "YES" || strings.ToUpper(choice) == "Y" {
				ValidLetter = true
			} else if strings.ToUpper(choice) != "NO" && strings.ToUpper(choice) != "N" {
				fmt.Println("\a\x1B[C\x1B[31manswer not valid\x1B[0m\n\x1B[C")
				time.Sleep(time.Second / 2)
			}
			ClearTerminal()

		case hangman.DoesContain(GameInProgress.RevealedLettres, guess):
			fmt.Printf("\a\x1B[C\x1B[31mThis letter has already been Revealed\x1B[0m\n\x1B[C")

		case hangman.DoesContain(GameInProgress.Guess, guess):
			fmt.Printf("\a\x1B[C\x1B[31mThis letter has already been chosen\x1B[0m\n\x1B[C")

		default:
			ValidLetter = true

		}
	}
	return strings.ToUpper(guess)
}
