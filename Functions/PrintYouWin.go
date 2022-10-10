package hangman

import (
	"fmt"
	"io/ioutil"
	"log"
)

func PrintYouWin() {
	content, err := ioutil.ReadFile("FilesAndLists/youwin.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
