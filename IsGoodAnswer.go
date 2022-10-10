package main

func IsGoodAnswer(GameInProgress Game, guess string) bool {
	for _, l := range GameInProgress.Word {
		if string(l) == guess || len(guess) > 1 && guess == GameInProgress.Word {
			return true
		}
	}
	return false
}
