package main

import (
	"fmt"
	hangman "hangman/Functions"
	"strings"
)

func StartPlaying(GameInProgress Game) { // go file in root to use structure "game" freely
	hangman.PrintWord(GameInProgress.Word, GameInProgress.RevealedLettres)

	for GameInProgress.Tries < 10 && !WordIsCompleted(GameInProgress) {
		fmt.Printf("*DEBUG* Revealed lettres are: %s\n", GameInProgress.RevealedLettres)
		fmt.Printf("*DEBUG* Already tried %s\n", GameInProgress.guess)
		guess := ChooseLetter(GameInProgress)
		GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, strings.ToUpper(guess))

		if IsGoodAnswer(GameInProgress, guess) {
			if len(guess) > 1 {
				for _, i := range guess {
					GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, string(i))
				}
			} else {
				GameInProgress.guess = append(GameInProgress.guess, guess)
				hangman.PrintWord(GameInProgress.Word, GameInProgress.RevealedLettres)
			}
		} else {
			if len(guess) > 1 && GameInProgress.Tries < 9 {
				GameInProgress.Tries++
			} else {
				GameInProgress.guess = append(GameInProgress.guess, guess)
			}
			fmt.Println("Try again!")
			PrintJose(GameInProgress)
			if GameInProgress.Tries < 9 {
				hangman.PrintWord(GameInProgress.Word, GameInProgress.RevealedLettres)
			}
			GameInProgress.Tries++
		}
	}
	if WordIsCompleted(GameInProgress) {
		hangman.PrintAscii("<YOU WIN>")
	}
}
