package main

import (
	"fmt"
)

func ChooseLetter(GameInProgress Game) string {
	var letter string
	fmt.Scanln(&letter)
	println("letter: ", letter)
	return letter
}
