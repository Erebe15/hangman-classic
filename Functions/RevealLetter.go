package hangman

import (
	"fmt"
	"math/rand"
	"time"
)

func RevealWord(s string) string {
	var RandomLetter []string
	rand.Seed(time.Now().UnixNano())
	random := s[rand.Intn(len(s))]
	if len(s) < 4 {
		RandomLetter = append(RandomLetter, string(random))
	} else {
		for i := 0; i < len(s)/2-1; i++ {
		RETURN:
			random := s[rand.Intn(len(s)-1)]
			if DoesContain(RandomLetter, string(random)) == true {
				goto RETURN
			} else {
				RandomLetter = append(RandomLetter, string(random))
			}
		}
	}
	for i := 0; i < len(s); i++ {
		if DoesContain(RandomLetter, string(s[i])) == true {
			fmt.Print(string(s[i]))
		} else {
			fmt.Print("_")
		}
	}
	return ""
}

func DoesContain(list []string, n string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == n {
			return true
		}
	}
	return false
}
