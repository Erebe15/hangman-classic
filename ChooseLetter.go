package main

import (
	"fmt"
	"strings"
)

func ChooseLetter(GameInProgress Game) Game {
	var letter string
	fmt.Scanln(&letter)
	println("letter: ", letter)
	GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, strings.ToUpper(letter))
	return GameInProgress
}
