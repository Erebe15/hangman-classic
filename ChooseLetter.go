package main

import (
	"fmt"
)

func ChooseLetter(GameInProgress Game) Game {
	var letter string
	fmt.Scan(&letter)
	println("letter: ", letter)
	GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, letter)
	return GameInProgress
	/*for i := 0; i < len(Word)-1; i++ {
		if DoesContain(starter, Word) == true {
			fmt.Print(string(Word[i]))
		} else {
			fmt.Print("_")

		}
	}
	*/
}
