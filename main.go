package main

import (
	"fmt"
	"hangman/Functions"
	"strings"
	"time"
)

type Game struct {
	Word                               string
	Tries                              int
	guess, RevealedLettres, JoseStates []string
}

type Window struct {
	ligns, colones int
}

var GameInProgress Game
var w Window

func Init() (bool, string) {
	w.ligns, w.colones = hangman.Size()
	GameInProgress.Tries = 0
	GameInProgress.JoseStates = hangman.GetJose()
	return true, ""
}

func welcome(l chan int) {
	hangman.Clear()
	time.Sleep(time.Second * 1)
	PrintAscii(3, w.colones/2-(len("<WELCOME>")*12)/2, "WELCOME")
	PrintAscii(12, w.colones/2-(len("<WELCOME>")*12)/2, "  TO   ")
	PrintAscii(21, w.colones/2-(len("<WELCOME>")*12)/2, "HANGMAN")
	time.Sleep(time.Second * 2)
	hangman.Clear()
	l <- 1
}

func main() {
	loading := make(chan int)
	go welcome(loading)
	PlayAgain, choice := Init()

	_ = <-loading
	for PlayAgain {

		CreateWindows()
		go resizeWindow()
		GameInProgress.Word = hangman.GetWord()
		fmt.Println(MoveTo(3, 2), "word:", GameInProgress.Word)
		GameInProgress.RevealedLettres = hangman.RevealStartLettres(GameInProgress.Word)
		StartPlaying()

		fmt.Println("")
		fmt.Println("Do you want to play again ?")
		fmt.Printf("Enter 'y' to play again, or any other input to quit : ")
		fmt.Scanln(&choice)
		if strings.ToUpper(choice) == "Y" || strings.ToUpper(choice) == "YES" {
			PlayAgain = true
		} else {
			fmt.Println("See you later!")
			PlayAgain = false
		}
	}
}
