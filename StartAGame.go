package main

import (
	hangman "hangman/Functions"
)

func WordIsCompleted() bool {
	for _, w := range GameInProgress.Word {
		if !hangman.DoesContain(GameInProgress.RevealedLettres, string(w)) {
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

func StartPlaying() {
	GameInProgress.Status = 1
	for GameInProgress.Tries < 10 && !WordIsCompleted() {

		ClearTerminal()
		UpdateWord()
		guess := ChooseLetter()

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
			if len(guess) > 1 && GameInProgress.Tries < 9 {
				GameInProgress.Tries++
			} else {
				GameInProgress.Guess = append(GameInProgress.Guess, guess)
			}
			WrongAnswEffect()
			PrintJose()
			GameInProgress.Tries++
		}
	}
	if WordIsCompleted() {
		GoodAnswerEffect()
		ClearAllWindows()
		GameInProgress.Status = 2
		PrintAscii(w.ligns/4-4, w.colones*65/200-42, "YOU WIN")
	} else {
		WrongAnswEffect()
		ClearAllWindows()
		GameInProgress.Status = 3
		PrintAscii(w.ligns/4-4, w.colones*65/200-42, "YOU LOST")
	}
}
