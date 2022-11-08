package main

import (
	"fmt"
	"hangman/Functions"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func IsAlpha(s string) bool {
	for _, l := range s {
		if l < 65 || l > 90 && l < 1040 || l > 1071 {
			return false
		}
	}
	return true
}

func ChooseLetter() string {
	var guess string
	ValidLetter := false
	var choice string

	for !ValidLetter {
		UpdateS()
		fmt.Scan(&guess)
		fmt.Print("\x1B[?25l")
		guess = strings.ToUpper(guess)

		switch {

		case guess == "STOP":
			ClearTerminal()
			docName := ""
			fmt.Print(GameInProgress.LanguageTxt[13], "\x1B[?25h") // Please enter the save name :
			fmt.Scanln(&docName)
			GameInProgress.Status += 10
			Marshal("Saves/" + docName + ".json")
			fmt.Printf("\x1B[?25l\x1B[C"+GameInProgress.LanguageTxt[14], docName) // Game saved as :
			time.Sleep(time.Second * 2)
			hangman.Clear()
			os.Exit(0)

		case guess == "DEVMODE":
			fmt.Print(MoveTo(3, 2), "the word is ", GameInProgress.Word)
			ClearTerminal()

		case !IsAlpha(guess):
			ClearTerminal()
			fmt.Print("\n\a\x1B[C\x1B[31m", GameInProgress.LanguageTxt[15], "\x1B[0m\n\x1B[C") // Please, select a letter

		case utf8.RuneCountInString(guess) > 1:
			ClearTerminal()
			fmt.Printf(GameInProgress.LanguageTxt[16]+"\n\x1B[C\x1B[32m"+GameInProgress.LanguageTxt[17]+"\x1B[0m / \x1B[31m"+GameInProgress.LanguageTxt[18]+"\x1B[0m\n\x1B[C\x1B[?25h", guess) // do you think it's %s? yes no
			fmt.Scanln(&choice)
			if strings.ToUpper(choice) == GameInProgress.LanguageTxt[19] || strings.ToUpper(choice) == GameInProgress.LanguageTxt[20] { // YES Y
				ValidLetter = true
			} else if strings.ToUpper(choice) != GameInProgress.LanguageTxt[21] && strings.ToUpper(choice) != GameInProgress.LanguageTxt[22] { // NO N
				fmt.Println("\a\x1B[C\x1B[31m", GameInProgress.LanguageTxt[23], "\x1B[0m\n\x1B[C") // answer not valid
				time.Sleep(time.Second)
			}
			ClearTerminal()

		case DoesContain(GameInProgress.RevealedLettres, guess):
			ClearTerminal()
			fmt.Print("\n\a\x1B[C\x1B[31m", GameInProgress.LanguageTxt[24], "\x1B[0m\n\x1B[C") // This letter has already been Revealed

		case DoesContain(GameInProgress.Guess, guess):
			ClearTerminal()
			fmt.Print("\n\a\x1B[C\x1B[31m", GameInProgress.LanguageTxt[25], "\x1B[0m\n\x1B[C") // This letter has already been chosen

		default:
			ValidLetter = true

		}
	}
	return strings.ToUpper(guess)
}
