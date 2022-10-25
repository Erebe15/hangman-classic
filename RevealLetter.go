package main

import (
	"math/rand"
	"time"
	"unicode/utf8"
)

func DoesContain(list []string, n string) bool {
	for _, l := range list {
		if l == n {
			return true
		}
	}
	return false
}

func RevealStartLettres(WordChosen string) []string { // select random lettres from the word for the start
	rand.Seed(time.Now().UnixNano())
	var RevealLetters []string
	var LettersOfWord []string
	var NumberOfRevealed int
	if GameInProgress.set.Difficulty >= 2 {
		NumberOfRevealed = 0
	} else {
		NumberOfRevealed = (utf8.RuneCountInString(WordChosen) / 2) - 1
	}

	for _, l := range WordChosen {
		LettersOfWord = append(LettersOfWord, string(l))
	}

	if utf8.RuneCountInString(WordChosen) >= 3 {
		for len(RevealLetters) < NumberOfRevealed {
			RndLetter := LettersOfWord[rand.Intn(len(LettersOfWord))]
			if !DoesContain(RevealLetters, RndLetter) {
				RevealLetters = append(RevealLetters, RndLetter)
			}
		}
	}

	return RevealLetters
}
