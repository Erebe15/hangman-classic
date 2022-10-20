package main

import (
	"flag"
	"fmt"
	hangman "hangman/Functions"
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
}

type Window struct {
	ligns, colones int
}

var GameInProgress Game
var w Window
var Save string

func init() {
	const (
		defaultSave = ""
		usage       = "the name of the save file"
	)
	flag.StringVar(&Save, "startWith", defaultSave, usage)
	flag.StringVar(&Save, "s", defaultSave, usage+" (shorthand)")
}

func gameInit(Ready chan int) (bool, string) {
	go resizeWindow(Ready)
	return true, ""
}

func NewGame() {
	GameInProgress.Word = hangman.GetWord()
	GameInProgress.Tries = 0
	GameInProgress.RevealedLettres = hangman.RevealStartLettres(GameInProgress.Word)
	GameInProgress.Status = 1
	GameInProgress.Guess = nil
	ClearAllWindows()
}

func main() {

	flag.Parse()
	if Save != "" {
		readJSON("Saves/" + Save + ".json")
		time.Sleep(time.Second * 3)
	} else {
		welcome()
	}

	Ready := make(chan int)
	PlayAgain, choice := gameInit(Ready)
	_ = <-Ready
	for PlayAgain {
		if GameInProgress.Status < 10 {
			NewGame()
		} else {
			GameInProgress.Status -= 10
		}

		StartPlaying()

		fmt.Scanln(&choice)
		fmt.Print("\x1B[?25l")
		if strings.ToUpper(choice) == "Y" || strings.ToUpper(choice) == "YES" || strings.ToUpper(choice) == "" {
			PlayAgain = true
		} else {
			PlayAgain = false
			fmt.Print("\x1B[Csee you later!")
			time.Sleep(time.Second)
			hangman.Clear()
		}
	}
}
