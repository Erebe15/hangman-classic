package main

func IsGoodAnswer(Inprogress Game, guess string) bool {
	for _, l := range Inprogress.Word {
		if string(l) == guess {
			return true
		}
	}
	return false
}
