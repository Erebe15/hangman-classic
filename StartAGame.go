package main

import (
	"fmt"
	"time"
)

func StartPlaying() {
	time.Sleep(time.Second)
	fmt.Print(MoveTo(w.ligns/2+2, 2))
	for GameInProgress.Tries < 10 && !WordIsCompleted(GameInProgress) {

		ClearTerminal()
		UpdateWord()
		guess := ChooseLetter(GameInProgress)

		if IsGoodAnswer(guess) {
			if len(guess) > 1 {
				for _, i := range guess {
					GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, string(i))
				}
			} else {
				GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, guess)
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
				PrintWord()
			}
			GameInProgress.Tries++
		}
	}
	if WordIsCompleted(GameInProgress) {
		CreateWindows()
		PrintAscii(w.ligns/4-4, w.colones*65/200-(9*12)/2, "YOU WIN")
	}
}
