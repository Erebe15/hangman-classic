package main

import (
	"fmt"
	"hangman/Functions"
)

func PrintWord() {
	display := ""
	for i := 0; i < len(GameInProgress.Word); i++ {
		if hangman.DoesContain(GameInProgress.RevealedLettres, string(GameInProgress.Word[i])) == true {
			display = display + string(GameInProgress.Word[i])
		} else {
			display = display + "_"
		}
	}
	PrintAscii(w.ligns/4-4, w.colones*65/200-(len(GameInProgress.Word)*12)/2, display)
	fmt.Print("\x1B[0m")
}
