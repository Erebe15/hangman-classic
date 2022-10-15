package main

import (
	"fmt"
	"os"
	"strings"
)

func PrintAscii(y, x int, word string) {
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
