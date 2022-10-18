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

func welcome() {
	w.ligns, w.colones = hangman.Size()
	hangman.Clear()
	PrintAscii(3, w.colones/2-(len("WELCOME")*12)/2, "WELCOME")
	PrintAscii(12, w.colones/2-(len("WELCOME")*12)/2, "  TO   ")
	PrintAscii(21, w.colones/2-(len("WELCOME")*12)/2, "HANGMAN")
	fmt.Print(MoveTo(0, 0))
	time.Sleep(time.Second * 2)
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
	for PlayAgain {
		if GameInProgress.Status < 10 {
			NewGame()
		} else {
			GameInProgress.Status -= 10
		}

		fmt.Println(MoveTo(3, 2), "word:", GameInProgress.Word)
		_ = <-Ready
		StartPlaying()

		ClearTerminal()
		fmt.Print("\x1B[CDo you want to play again ?\n\x1B[C Enter 'y' to play again, or any other input to quit : ")
		fmt.Scanln(&choice)
		if strings.ToUpper(choice) == "Y" || strings.ToUpper(choice) == "YES" {
			PlayAgain = true
		} else {
			PlayAgain = false
			fmt.Print("\x1B[Csee you later!")
			time.Sleep(1 * time.Second)
			hangman.Clear()
		}
	}
}
