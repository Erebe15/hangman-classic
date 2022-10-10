package hangman

import (
	"log"
	"os"
	"strings"
)

func GetJose() []string {
	data, err := os.ReadFile("FilesAndLists/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	var JosePositions []string
	JosePositions = strings.Split(string(data), "♦")
	return JosePositions
}
