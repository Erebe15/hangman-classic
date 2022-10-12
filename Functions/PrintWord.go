package hangman

import "fmt"

func PrintWord(Word string, LetterList []string) {
	display := ""
	for i := 0; i < len(Word); i++ {
		if DoesContain(LetterList, string(Word[i])) == true {
			//fmt.Printf(" %s", string(Word[i]))
			display = display + string(Word[i])
		} else {
			display = display + "_"
		}
	}
	PrintAscii("<" + display + ">")
	fmt.Printf("\n")
}
