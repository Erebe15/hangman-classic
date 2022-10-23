package main

import (
	"fmt"
	hangman "hangman/Functions"
)

var Esc = "\x1b"

func escape(format string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", Esc, fmt.Sprintf(format, args...))
}

func MoveTo(line, col int) string {
	return escape("[%d;%dH", line, col) // "[8;15H"
}

func cadre(y1, x1, y2, x2 int, title string) {
	fmt.Print(MoveTo(y1, x1), "╔")
	fmt.Print(MoveTo(y1, x2), "╗")
	fmt.Print(MoveTo(y2, x1), "╚")
	fmt.Print(MoveTo(y2, x2), "╝")
	for i := x1 + 1; i < x2; i++ {
		fmt.Print(MoveTo(y1, i), "═")
		fmt.Print(MoveTo(y2, i), "═")
	}
	for i := y1 + 1; i < y2; i++ {
		fmt.Print(MoveTo(i, x1), "║")
		fmt.Print(MoveTo(i, x2), "║")
	}
	fmt.Print(MoveTo(y1, (x1+x2)/2-(len(title)/2)), title)
	fmt.Print(MoveTo(w.ligns/2+2, 2), "")
}

func CreateWindows() {
	cadre(1, 1, w.ligns/2, w.colones*65/100, GameInProgress.set.LanguageTxt[2])           // WORD
	cadre(1, (w.colones*65/100)+1, w.ligns, w.colones, GameInProgress.set.LanguageTxt[4]) // HANGMAN
	cadre(w.ligns/2+1, 1, w.ligns, w.colones*65/100, GameInProgress.set.LanguageTxt[3])   // TERMINAL
	fmt.Print(MoveTo(w.ligns/2+2, 2))
}

func resizeWindow(ready chan int) {
	hangman.Clear()
	CreateWindows()
	ready <- 1
	for {
		a, b := hangman.Size()
		if w.ligns != a || w.colones != b {
			hangman.Clear()
			w.ligns, w.colones = hangman.Size()
			fmt.Print("\x1B[0m")
			CreateWindows()
			UpdateS()
		}
	}
}

func ClearTerminal() {
	for i := w.ligns/2 + 2; i < w.ligns; i++ {
		for j := 2; j < w.colones*65/100; j++ {
			fmt.Print(MoveTo(i, j), " ")
		}
	}
	fmt.Print(MoveTo(w.ligns/2+2, 2))
}

func ClearAllWindows() {
	for i := w.ligns/2 + 2; i < w.ligns; i++ {
		for j := 2; j < w.colones*65/100; j++ {
			fmt.Print(MoveTo(i, j), " ")
		}
	}
	for i := 2; i < w.ligns/2; i++ {
		for j := 2; j < w.colones*65/100; j++ {
			fmt.Print(MoveTo(i, j), " ")
		}
	}
	for i := 2; i < w.ligns; i++ {
		for j := w.colones*65/100 + 2; j < w.colones; j++ {
			fmt.Print(MoveTo(i, j), " ")
		}
	}
	fmt.Print(MoveTo(w.ligns/2+2, 2))
}

func UpdateS() {
	switch GameInProgress.Status {
	case 1:
		PrintWord()
		DrawHangmanState()
		if GameInProgress.RevealedLettres != nil {
			fmt.Print(MoveTo(2, 2), "\x1B[32m", GameInProgress.set.LanguageTxt[5], GameInProgress.RevealedLettres) // REVEALED LETTERS :
		}
		if GameInProgress.Guess != nil {
			fmt.Print(MoveTo(2, w.colones*65/100+2), "\x1B[31m", GameInProgress.set.LanguageTxt[6], GameInProgress.Guess) // WRONG GUESS :
		}
		fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[34m", GameInProgress.set.LanguageTxt[7], "\x1B[0m\x1B[?25h") // Choose a letter:
	case 2:
		ClearAllWindows()
		PrintAscii(w.ligns/4-3, w.colones*65/200-44, GameInProgress.set.LanguageTxt[8]) // YOU WIN
		fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[C", GameInProgress.set.LanguageTxt[10], "\x1B[32m", GameInProgress.Word, "\x1B[0m! ", GameInProgress.set.LanguageTxt[11], "\n\x1B[C ", GameInProgress.set.LanguageTxt[12], " \x1B[?25h")
	case 3:
		ClearAllWindows()
		fmt.Print("\x1B[31m")
		PrintAscii(w.ligns/4-3, w.colones*65/200-46, GameInProgress.set.LanguageTxt[9]) // YOU LOST
		DrawHangmanState()
		fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[C", GameInProgress.set.LanguageTxt[10], " \x1B[31m", GameInProgress.Word, "\x1B[0m! ", GameInProgress.set.LanguageTxt[11], "\n\x1B[C ", GameInProgress.set.LanguageTxt[12], " \x1B[?25h")
	}
}
