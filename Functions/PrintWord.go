package hangman

import "fmt"

func PrintWord(Word string, LetterList []string) {
	fmt.Printf("liste de lettre revealed: %s\n", LetterList)
	for i := 0; i < len(Word); i++ {
		if DoesContain(LetterList, string(Word[i])) == true {
			fmt.Print(string(Word[i]))
		} else {
			fmt.Print("_")
		}
	}
	fmt.Print("\n")
}
