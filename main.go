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

1 In game
2 Won
3 Lost
11 Saved playing
12 Saved won
13 Saved lost

*/

type Game struct {
	Word                   string
	Tries, Status          int
	Guess, RevealedLettres []string
	set                    Settings
}

type Settings struct {
	LanguageTxt []string
	Words       string
	Difficulty  int
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
	GameInProgress.set.LanguageTxt = strings.Split(string(data), string(sep))
	switch language {
	case "English":
		GameInProgress.set.Words = "words.txt"
	case "Francais":
		GameInProgress.set.Words = "mots.txt"
	case "Russian":
		GameInProgress.set.Words = "слова.txt"
	}
}

func gameInit(Ready chan int) (bool, string) {
	go resizeWindow(Ready)
	return true, ""
}

func NewGame() {
	GameInProgress.Word = GetWord()
	GameInProgress.Tries = 0
	GameInProgress.RevealedLettres = hangman.RevealStartLettres(GameInProgress.Word)
	GameInProgress.Status = 1
	GameInProgress.Guess = nil
	GameInProgress.set.Difficulty = 1
	ClearAllWindows()
}

func main() {
	w.ligns, w.colones = hangman.Size()
	flag.Parse()
	if Save != "" {
		readJSON("Saves/" + Save + ".json")
		time.Sleep(time.Second * 3)
	} else if w.colones >= 155 && w.ligns >= 41 {
		WelcomeArrive()
		MainMenu()
		WelcomeFadeAway()
	}

	WindowsReady := make(chan int)
	PlayAgain, choice := gameInit(WindowsReady)
	_ = <-WindowsReady

	for PlayAgain {
		if GameInProgress.Status < 10 {
			NewGame()
		} else {
			GameInProgress.Status -= 10
		}

		StartPlaying()

		fmt.Scanln(&choice)
		fmt.Print("\x1B[?25l")
		if strings.ToUpper(choice) == GameInProgress.set.LanguageTxt[19] || strings.ToUpper(choice) == GameInProgress.set.LanguageTxt[20] || strings.ToUpper(choice) == "" { // YES Y
			PlayAgain = true
		} else {
			PlayAgain = false
			fmt.Print("\x1B[C", GameInProgress.set.LanguageTxt[26]) // see you later
			time.Sleep(time.Second)
			hangman.Clear()
		}
	}
}
