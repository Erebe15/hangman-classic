package hangman

import (
	"fmt"
	"io/ioutil"
	"log"
)

func PrintYouWin() {
	content, err := ioutil.ReadFile("WordsLists/youwin.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
