package hangman

import (
	"fmt"
	"os"
	"strings"
)

func PrintAscii(word string) {
	data, err := os.ReadFile("FilesAndLists/fancyLetters.txt")
	if err != nil {
		panic(err)
	}
	var ligns [][]string
	sep := []byte{13, 10, 13, 10}
	char := strings.Split(string(data), string(sep))
	for _, l := range char {
		ligns = append(ligns, strings.Split(l, string(sep[2:])))
	}
	characters := "<> -();,._ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for floor := 0; floor < 9; floor++ {
		for _, l := range word {
			for i, c := range characters {
				if l == c {
					fmt.Print(ligns[i][floor])
				}
			}
		}
		fmt.Printf("\n")
	}
}
