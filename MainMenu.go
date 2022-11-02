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

func Menu4(a, b, c, d int, fa, fb, fc, fd, fe func()) {
	var choice string
	ClearMenu()
	for i := 232; i <= 250; i++ {
		fmt.Printf("\x1B[38;5;%dm", i)
		fmt.Print(MoveTo(w.ligns-1, w.colones-utf8.RuneCountInString("(10) "+hangman.NicerTypo(GameInProgress.LanguageTxt[45]))), "(10) "+hangman.NicerTypo(GameInProgress.LanguageTxt[45]))
		fmt.Print(MoveTo(33, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[a])/2), GameInProgress.LanguageTxt[a])
		fmt.Print(MoveTo(36, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[b])/2), GameInProgress.LanguageTxt[b])
		fmt.Print(MoveTo(39, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[c])/2), GameInProgress.LanguageTxt[c])
		fmt.Print(MoveTo(42, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[d])/2), GameInProgress.LanguageTxt[d])
		time.Sleep(time.Millisecond * 20)
	}
	fmt.Print("\x1B[0m")
	fmt.Print(MoveTo(w.ligns-1, w.colones-utf8.RuneCountInString("(10) "+hangman.NicerTypo(GameInProgress.LanguageTxt[45]))), "(10) "+hangman.NicerTypo(GameInProgress.LanguageTxt[45]))
	fmt.Print(MoveTo(33, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[a])/2), GameInProgress.LanguageTxt[a])
	fmt.Print(MoveTo(36, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[b])/2), GameInProgress.LanguageTxt[b])
	fmt.Print(MoveTo(39, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[c])/2), GameInProgress.LanguageTxt[c])
	fmt.Print(MoveTo(42, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[d])/2), GameInProgress.LanguageTxt[d])
	fmt.Print(MoveTo(1, 1), ">_")

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
	case choice == "10":
		MainMenu()
	default:
		fe()
	}
}

func LanguageMenu() {
	Menu4(39, 40, 41, 34, LanguageMenuEn, LanguageMenuFr, LanguageMenuRu, LanguageMenu, OptionsMenu)
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
	if GameInProgress.Words != "mots.txt" {
		SelectLanguage("Francais")
		hangman.Clear()
		WelcomeArrive()
	}
	MainMenu()
}

func LanguageMenuEn() {
	if GameInProgress.Words != "words.txt" {
		SelectLanguage("English")
		hangman.Clear()
		WelcomeArrive()
	}
	MainMenu()
}

func LanguageMenuRu() {
	if GameInProgress.Words != "слова.txt" {
		SelectLanguage("Russian")
		hangman.Clear()
		WelcomeArrive()
	}
	MainMenu()
}

func DifficultySet1() {
	GameInProgress.Difficulty = 1
	MainMenu()
}

func DifficultySet2() {
	GameInProgress.Difficulty = 2
	MainMenu()
}

func DifficultySet3() {
	GameInProgress.Difficulty = 3
	MainMenu()
}

func DifficultySet4() {
	GameInProgress.Difficulty = 4
	MainMenu()
}

func Return() {
	return
}

func quit() {
	hangman.Clear()
	os.Exit(0)
}
