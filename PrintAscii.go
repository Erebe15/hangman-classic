package main

import (
	"fmt"
	hangman "hangman/Functions"
	"os"
	"strings"
)

func PrintWord() {
	display := ""
	for i := 0; i < len(GameInProgress.Word); i++ {
		if hangman.DoesContain(GameInProgress.RevealedLettres, string(GameInProgress.Word[i])) == true {
			if len(GameInProgress.Word)*12 > w.colones*65/100-4 || w.ligns/2 < 11 {
				display += " " + string(GameInProgress.Word[i])
			} else {
				display = display + string(GameInProgress.Word[i])
			}
		} else {
			if len(GameInProgress.Word)*12 > w.colones*65/100-4 || w.ligns/2 < 11 {
				display += " _"
			} else {
				display += "_"
			}
		}
	}
	PrintAscii(w.ligns/4-3, w.colones*65/200-(len(GameInProgress.Word)*12)/2, display)
	fmt.Print("\x1B[0m")
}

func PrintAscii(y, x int, word string) {
	if len(word)*12 > w.colones*65/100-4 || w.ligns/2 < 11 {
		fmt.Print(MoveTo(y+3, x+len(word)*3-len(word)/2), "\x1B[1m"+word)
		fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[0m")
		return
	}
	fmt.Print("\x1B[7m")
	data, err := os.ReadFile("FilesAndLists/fancyLetters.txt")
	if err != nil {
		panic(err)
	}
	var CharLigns [][]string
	sep := []byte{13, 10, 13, 10}
	char := strings.Split(string(data), string(sep))
	for _, l := range char {
		CharLigns = append(CharLigns, strings.Split(l, string(sep[2:])))
	}
	characters := "<> -();,._ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for floor := 0; floor < 9; floor++ {
		for k, l := range word {
			for i, c := range characters {
				if l == c {
					fmt.Print(MoveTo(y+floor, x+k*12), CharLigns[i][floor])
				}
			}
		}
	}
	fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[0m")
}
