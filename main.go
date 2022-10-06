package main

import "fmt"

type hangman struct {
	Advancment, Word string
	tried            int
}

func main() {
	hangman.Word = hangman.GetWord()
	fmt.Println(hangman.word)
}
