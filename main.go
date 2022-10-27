package main

import (
	"flag"
	"fmt"
	hangman "hangman/Functions"
	"os"
	"strings"
	"time"
)

/*

Game Status:

0 Menu
1 In game
2 Won
3 Lost
11 Saved playing
12 Saved won
13 Saved lost

*/

type Game struct {
	Word, Words                         string
	Tries, Status, Difficulty           int
	Guess, RevealedLettres, LanguageTxt []string
}

type Window struct {
	ligns, colones int
}

var GameInProgress Game
var w Window
var Save string

func init() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "mots.txt":
			SelectLanguage("Francais")
		case "слова.txt":
			SelectLanguage("Russian")
		default:
			SelectLanguage("English")
		}
	} else {
		SelectLanguage("English")
	}

	const (
		defaultSave = ""
		usage       = "the name of the save file"
	)
	flag.StringVar(&Save, "startWith", defaultSave, usage)
	flag.StringVar(&Save, "s", defaultSave, usage+" (shorthand)")
}

func SelectLanguage(language string) {
	data, err := os.ReadFile("FilesAndLists/Languages/" + language + ".txt")
	if err != nil {
		panic(err)
	}
	sep := []byte{13, 10}
	GameInProgress.LanguageTxt = strings.Split(string(data), string(sep))
	switch language {
	case "English":
		GameInProgress.Words = "words.txt"
	case "Francais":
		GameInProgress.Words = "mots.txt"
	case "Russian":
		GameInProgress.Words = "слова.txt"
	}
}

func NewGame() {
	GameInProgress.Word = GetWord()
	GameInProgress.Tries = 0
	GameInProgress.RevealedLettres = RevealStartLettres(GameInProgress.Word)
	GameInProgress.Status = 1
	GameInProgress.Guess = nil
	ClearAllWindows()
}

func LaunchMenu() {
	WelcomeArrive()
	MainMenu()
	WelcomeFadeAway()
}

func main() {
	w.ligns, w.colones = hangman.Size()
	flag.Parse()
	if Save != "" {
		readJSON("Saves/" + Save + ".json")
		time.Sleep(time.Second * 3)
	} else if w.colones >= 155 && w.ligns >= 41 {
		LaunchMenu()
	}

	WindowsReady := make(chan int)
	go resizeWindow(WindowsReady)
	PlayAgain := true
	var choice string
	_ = <-WindowsReady

	for PlayAgain {
		if GameInProgress.Status < 10 { // !Saved game
			NewGame()
		}

		StartPlaying()

		for choice != "1" && choice != "2" && choice != "3" {
			fmt.Scanln(&choice)
			fmt.Print("\x1B[?25l")
			if choice == "2" {
				PlayAgain = true
				GameInProgress.Status = 0 // in the menu
				LaunchMenu()
				GameInProgress.Status = 1 // in game
				hangman.Clear()
				CreateWindows()
			} else if choice == "3" {
				PlayAgain = false
				fmt.Print("\x1B[C", GameInProgress.LanguageTxt[26]) // see you later
				time.Sleep(time.Second)
				hangman.Clear()
			}
		}
	}
}
