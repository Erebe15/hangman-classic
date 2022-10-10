package main

import hangman "hangman/Functions"

func WordIsCompleted(InProgress Game) bool {
	for _, w := range InProgress.Word {
		if !hangman.DoesContain(InProgress.RevealedLettres, string(w)) {
			return false
		}
	}
	return true
}
