package main

import hangman "hangman/Functions"

func StartPlaying(GameInProgress Game) { // go file in root to use struc "game" freely
	print("tries : ")
	hangman.PrintWord(GameInProgress.Word, GameInProgress.RevealedLettres)
	for GameInProgress.Tries < 10 && !WordCompleted(GameInProgress) {
		// print word
		//ask for input and ad it to the list of guess
		print(GameInProgress.Tries, ", ")
		GameInProgress.Tries++
	}
	println(GameInProgress.Tries)
	println("ggwp...")
}
