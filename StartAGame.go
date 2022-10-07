package main

import (
	"fmt"
	hangman "hangman/Functions"
)

func StartPlaying(GameInProgress Game) { // go file in root to use struc "game" freely

	for GameInProgress.Tries < 10 && !WordCompleted(GameInProgress) {
		hangman.PrintWord(GameInProgress.Word, GameInProgress.RevealedLettres)
		fmt.Printf("Choose a letter : ")
		GameInProgress = ChooseLetter(GameInProgress)
		fmt.Printf("Revealed lettres are: %s\n", GameInProgress.RevealedLettres)
		GameInProgress.Tries++
	}

	println("ggwp...")
}
