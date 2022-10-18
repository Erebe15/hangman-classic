package main

import (
	"fmt"
	hangman "hangman/Functions"
	"time"
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
	cadre(1, 1, w.ligns/2, w.colones*65/100, "WORD")
	cadre(1, (w.colones*65/100)+1, w.ligns, w.colones, "HANGMAN")
	cadre(w.ligns/2+1, 1, w.ligns, w.colones*65/100, "TERMINAL")
	fmt.Print(MoveTo(w.ligns/2+2, 2))
}

func WrongAnswEffect() {
	for i := 0; i <= 6; i++ {
		if i%2 == 1 {
			fmt.Print("\x1B[31m")
		} else {
			fmt.Print("\x1B[0m")
		}
		CreateWindows()
		time.Sleep(time.Second / 10)
	}
	fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[0m")
}

func GoodAnswerEffect() {
	fmt.Print("\x1B[32m")
	CreateWindows()
	fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[0m")
	time.Sleep(time.Second / 3)
	CreateWindows()
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
		for j := w.colones*65/100 + 1; j < w.colones*65/100; j++ {
			fmt.Print(MoveTo(i, j), " ")
		}
	}
	fmt.Print(MoveTo(w.ligns/2+2, 2))
}

func UpdateWord() {
	switch GameInProgress.Status {
	case 1:
		PrintWord()
		if GameInProgress.RevealedLettres != nil {
			fmt.Print(MoveTo(2, 2), "\x1B[32mREVEALED LETTERS : ", GameInProgress.RevealedLettres)
		}
	case 2:
		ClearAllWindows()
		PrintAscii(w.ligns/4-4, w.colones*65/200-42, "YOU WIN")
	case 3:
		ClearAllWindows()
		PrintAscii(w.ligns/4-4, w.colones*65/200-42, "YOU LOST")
	}
	fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[0m")
}
