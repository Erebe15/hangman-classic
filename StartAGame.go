package main

func StartPlaying(GameInProgress Game) {
	for GameInProgress.Tries <= 10 && !WordCompleted(GameInProgress) {
		// print word
		//ask for input and ad it to the list of guess
		GameInProgress.Tries++
	}
}
