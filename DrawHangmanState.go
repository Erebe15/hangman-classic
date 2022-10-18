package main

import (
	"fmt"
	"os"
	"strings"
)

func PrintJose() {
	data, err := os.ReadFile("FilesAndLists/hangman.txt")
	if err != nil {
		panic(err)
	}
	var JoseFloor [][]string
	States := strings.Split(string(data), "\n\n")
	for _, i := range States {
		JoseFloor = append(JoseFloor, strings.Split(i, "\n"))
	}
	for j, l := range JoseFloor[GameInProgress.Tries] {
		fmt.Print(MoveTo(w.ligns/2-4+j, w.colones*75/100), l)
	}
}

func DrawHangmanState() {
	if w.ligns >= 10000 && w.colones > 7 {
		drawFancy()
	} else {
		PrintJose()
	}
}

func drawFancy() {
	if GameInProgress.Tries >= 1 {
	}
	if GameInProgress.Tries >= 2 {
	}
	if GameInProgress.Tries >= 3 {
	}
	if GameInProgress.Tries >= 4 {
	}
	if GameInProgress.Tries >= 5 {
	}
	if GameInProgress.Tries >= 6 {
	}
	if GameInProgress.Tries >= 7 {
	}
	if GameInProgress.Tries >= 8 {
	}
	if GameInProgress.Tries >= 9 {
	}
	if GameInProgress.Tries >= 10 {
	}
}
