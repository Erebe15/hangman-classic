package main

import (
	"fmt"
)

func WordIsCompleted() bool {
	for _, w := range GameInProgress.Word {
		if !DoesContain(GameInProgress.RevealedLettres, string(w)) {
			return false
		}
	}
	return true
}

func IsGoodAnswer(guess string) bool {
	for _, l := range GameInProgress.Word {
		if string(l) == guess || len(guess) > 1 && guess == GameInProgress.Word {
			return true
		}
	}
	return false
}

func ProcessGuess(guess string) {
	if IsGoodAnswer(guess) {
		if len(guess) > 1 {
			for _, i := range guess {
				GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, string(i))
			}
		} else {
			GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, guess)
			GoodAnswerEffect()
		}
	} else {
		WrongAnswerEffect()
		if len(guess) > 1 && GameInProgress.Tries < 9 {
			GameInProgress.Tries++
			HangmanAnimation()
		} else {
			GameInProgress.Guess = append(GameInProgress.Guess, guess)
		}
		GameInProgress.Tries++
		HangmanAnimation()
	}
}

func StartPlaying() {
	GameInProgress.Status = 1 // playing
	fmt.Print(MoveTo(3, 2), "the word: ", GameInProgress.Word)
	for GameInProgress.Tries < 10 && !WordIsCompleted() {

		ClearTerminal()
		guess := ChooseLetter()
		ProcessGuess(guess)
	}

	if WordIsCompleted() {
		GoodAnswerEffect()
		ClearAllWindows()
		GameInProgress.Status = 2 // won game
	} else {
		GameInProgress.Status = 3 // lost game
	}
	UpdateS()
}
