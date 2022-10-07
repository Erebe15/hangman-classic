package hangman

import (
	"log"
	"os"
	"runtime"
)

func GetJose() []string {
	data, err := os.ReadFile("WordsLists/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	os := runtime.GOOS //runtime.GOOS -> linux, windows, darwin etc.
	println("\nYou are running on", os, "!")

	var JosePositions []string
	var char string
	for i, l := range data {
		if string(l) == "\n" && string(data[i-1]) == "\n" {
			if os == "windows" {
				JosePositions = append(JosePositions, char[:len(char)-2]) // -1 getting rid of the backslash (byte '13')
			} else {
				JosePositions = append(JosePositions, char[:len(char)])
			}
			char = ""
		} else {
			char = char + string(l)
		}
	}
	return JosePositions
}
