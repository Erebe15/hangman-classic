package main

import (
	hangman "hangman/Functions"
	"strings"
)

func StartPlaying(GameInProgress Game) { // go file in root to use structure "game" freely

	hangman.PrintWord(GameInProgress.Word, GameInProgress.RevealedLettres)
	for GameInProgress.Tries < 10 && !WordIsCompleted(GameInProgress) {
		guess := ChooseLetter(GameInProgress)
		GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, strings.ToUpper(guess))
		if IsGoodAnswer(GameInProgress, guess) {
			hangman.PrintWord(GameInProgress.Word, GameInProgress.RevealedLettres)
		} else {
			PrintJose(GameInProgress)
			GameInProgress.Tries++
		}
	}
	if WordIsCompleted(GameInProgress) {
		println("_______GG WP_______")

	} else {
		println("_____GAME OVER_____")
	}

}
