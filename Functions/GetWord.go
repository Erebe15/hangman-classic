package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

func GetWord() string {
	if len(os.Args) == 1 {
		fmt.Println("Please select a file as argument")
		os.Exit(1)
	}
	data, err := os.ReadFile(string(os.Args[1]))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Words := strings.Split(string(data), "\n")
	os := runtime.GOOS //runtime.GOOS -> linux, windows, darwin etc.
	if os == "windows" {
		for i, word := range Words {
			Words[i] = word[:len(word)-1]
		}
	}
	rand.Seed(time.Now().UnixNano())
	word := strings.ToUpper(Words[rand.Intn(len(Words))])
	return word
}
