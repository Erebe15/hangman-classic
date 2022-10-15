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
}

func CreateWindows() {
	hangman.Clear()
	w.ligns, w.colones = hangman.Size()
	cadre(1, 1, w.ligns/2, w.colones*65/100, "WORD")
	cadre(1, (w.colones*65/100)+1, w.ligns, w.colones, "HANGMAN")
	cadre(w.ligns/2+1, 1, w.ligns, w.colones*65/100, "TERMINAL")
	fmt.Print(MoveTo(w.ligns/2+2, 2))
}

func resizeWindow() {
	for {
		a, b := hangman.Size()
		if w.ligns != a || w.colones != b {
			CreateWindows()
			UpdateWord()
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

func UpdateWord() {
	PrintWord()
	fmt.Print(MoveTo(2, 2), "\x1B[7mREVEALED LETTERS : ", GameInProgress.RevealedLettres)
	fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[0m")
}
