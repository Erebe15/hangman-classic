package hangman

import (
	"fmt"
	"os"
)

func PrintRules() {
	data, err := os.ReadFile("FilesAndLists/Welcome.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
