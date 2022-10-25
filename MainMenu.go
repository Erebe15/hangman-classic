package main

import (
	"fmt"
	hangman "hangman/Functions"
	"os"
	"time"
	"unicode/utf8"
)

func ClearMenu() {
	fmt.Print(MoveTo(1, 1), "     ")
	fmt.Print(MoveTo(33, w.colones/2-30))
	for l := 0; l <= 10; l++ {
		fmt.Print("                                                            \x1B[B\x1B[60D")
	}
}

func PrintMenu() {
	for i := 232; i <= 250; i++ {
		fmt.Printf("\x1B[38;5;%dm", i)
		fmt.Print(MoveTo(33, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[27])/2), GameInProgress.set.LanguageTxt[27]) // N E W   G A M E
		fmt.Print(MoveTo(36, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[27])/2), GameInProgress.set.LanguageTxt[28]) // S A V E S
		fmt.Print(MoveTo(39, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[27])/2), GameInProgress.set.LanguageTxt[29]) // O P T I O N S
		fmt.Print(MoveTo(42, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[27])/2), GameInProgress.set.LanguageTxt[30]) // Q U I T
		time.Sleep(time.Millisecond * 20)
	}
	fmt.Print("\x1B[0m")
	fmt.Print(MoveTo(33, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[27])/2), GameInProgress.set.LanguageTxt[27])
	fmt.Print(MoveTo(36, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[27])/2), GameInProgress.set.LanguageTxt[28])
	fmt.Print(MoveTo(39, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[27])/2), GameInProgress.set.LanguageTxt[29])
	fmt.Print(MoveTo(42, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[27])/2), GameInProgress.set.LanguageTxt[30])
	fmt.Print(MoveTo(1, 1), ">_")
}

func Menu4(a, b, c, d int, fa, fb, fc, fd, fe func()) {
	var choice string
	ClearMenu()
	for i := 232; i <= 250; i++ {
		fmt.Printf("\x1B[38;5;%dm", i)
		fmt.Print(MoveTo(33, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[a])/2), GameInProgress.set.LanguageTxt[a])
		fmt.Print(MoveTo(36, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[b])/2), GameInProgress.set.LanguageTxt[b])
		fmt.Print(MoveTo(39, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[c])/2), GameInProgress.set.LanguageTxt[c])
		fmt.Print(MoveTo(42, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[d])/2), GameInProgress.set.LanguageTxt[d])
		time.Sleep(time.Millisecond * 20)
	}
	fmt.Print("\x1B[0m")
	fmt.Print(MoveTo(33, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[a])/2), GameInProgress.set.LanguageTxt[a])
	fmt.Print(MoveTo(36, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[b])/2), GameInProgress.set.LanguageTxt[b])
	fmt.Print(MoveTo(39, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[c])/2), GameInProgress.set.LanguageTxt[c])
	fmt.Print(MoveTo(42, w.colones/2-utf8.RuneCountInString(GameInProgress.set.LanguageTxt[d])/2), GameInProgress.set.LanguageTxt[d])
	fmt.Print(MoveTo(1, 1), ">_")

	ValidInput := []string{"1", "2", "3", "4"}
	fmt.Scanln(&choice)
	switch {
	case choice == "1":
		fa()
	case choice == "2":
		fb()
	case choice == "3":
		fc()
	case choice == "4":
		fd()
	case !hangman.DoesContain(ValidInput, choice):
		fe()
	}
}

func LanguageMenu() {
	Menu4(39, 40, 41, 33, LanguageMenuEn, LanguageMenuFr, LanguageMenuRu, OptionsMenu, LanguageMenu)
}

func DifficultyMenu() {
	Menu4(35, 36, 37, 38, DifficultySet1, DifficultySet2, DifficultySet3, DifficultySet4, OptionsMenu)
}

func OptionsMenu() {
	Menu4(31, 32, 33, 34, LanguageMenu, DifficultyMenu, MainMenu, OptionsMenu, OptionsMenu)
}

func MainMenu() {
	Menu4(27, 28, 29, 30, Return, Saves, OptionsMenu, quit, MainMenu)
}

func LanguageMenuFr() {
	if GameInProgress.set.Words != "mots.txt" {
		SelectLanguage("Francais")
		hangman.Clear()
		WelcomeArrive()
	}
	MainMenu()
}

func LanguageMenuEn() {
	if GameInProgress.set.Words != "words.txt" {
		SelectLanguage("English")
		hangman.Clear()
		WelcomeArrive()
	}
	MainMenu()
}

func LanguageMenuRu() {
	if GameInProgress.set.Words != "слова.txt" {
		SelectLanguage("Russian")
		hangman.Clear()
		WelcomeArrive()
	}
	MainMenu()
}

func DifficultySet1() {
	GameInProgress.set.Difficulty = 1
	OptionsMenu()
}

func DifficultySet2() {
	GameInProgress.set.Difficulty = 2
	OptionsMenu()
}

func DifficultySet3() {
	GameInProgress.set.Difficulty = 3
	OptionsMenu()
}

func DifficultySet4() {
	GameInProgress.set.Difficulty = 4
	OptionsMenu()
}

func Return() {
	return
}

func quit() {
	os.Exit(0)
}
