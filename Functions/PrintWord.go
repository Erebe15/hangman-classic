package hangman

import "fmt"

func PrintWord(Word string, LetterList []string) {
	for i := 0; i < len(Word); i++ {
		if DoesContain(LetterList, string(Word[i])) == true {
			fmt.Print(string(Word[i]))
		} else {
			fmt.Print("_")
		}
	}
	fmt.Print("\n")
}
