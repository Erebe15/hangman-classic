package main

import (
	"fmt"
	"hangman/Functions"
	"os"
	"os/exec"
	"strings"
)

type Game struct {
	Word                                     string
	Tries                                    int
	guess, RevealedLettres, JoseStates, save []string
}

func (Game) init() Game {
	var GameInProgress Game
	GameInProgress.Word = hangman.GetWord()
	GameInProgress.RevealedLettres = hangman.RevealStartLettres(GameInProgress.Word)
	GameInProgress.Tries = 0
	GameInProgress.JoseStates = hangman.GetJose()
	return GameInProgress
}

func main() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	var GameInProgress Game
	PlayAgain := true
	choice := ""

	if os.Args[1] == "--startwith" {
		GameInProgress.readJSON(os.Args[2] + ".json")

	} else {
		GameInProgress = GameInProgress.init()
	}
	for PlayAgain {

		hangman.PrintRules()
		StartPlaying(GameInProgress)

		fmt.Println("")
		fmt.Println("Do you want to play again ?")
		fmt.Printf("Enter 'y' to play again, or any other input to quit : ")
		fmt.Scanln(&choice)
		if strings.ToUpper(choice) == "Y" {
			GameInProgress = GameInProgress.init()
			PlayAgain = true
		} else {
			fmt.Println("See you later!")
			PlayAgain = false
		}
	}
}
