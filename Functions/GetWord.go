package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetWord() string {
	rand.Seed(time.Now().UnixNano())
	sep := []byte{13, 10}
	data, err := os.ReadFile("FilesAndLists/words.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		data, _ = os.ReadFile("FilesAndLists/" + string(os.Args[1]))
	}
	Words := strings.Split(string(data), string(sep))
	word := strings.ToUpper(Words[rand.Intn(len(Words))])
	return word
}
