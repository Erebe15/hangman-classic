package main

import (
	"encoding/json"
	"fmt"
	hangman "hangman/Functions"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"time"
	"unicode/utf8"
)

func Marshal(file string) {
	b, err := json.Marshal(GameInProgress)
	if err != nil {
		fmt.Println("error", err)
	}
	os.WriteFile(file, b, 0644)
}

func readJSON(file string) {
	jsonFile, err := os.Open(file)
	err = json.NewDecoder(jsonFile).Decode(&GameInProgress)
	err = jsonFile.Close()
	err = os.Remove(file)
	if err != nil {
		fmt.Println("error = ", err.Error())
	}
}

func Saves() {
	SavesFiles, err := ioutil.ReadDir("Saves/")
	if err != nil {
		log.Fatal(err)
	}
	if len(SavesFiles) == 0 {
		ClearMenu()
		fmt.Print(MoveTo(33, w.colones/2-utf8.RuneCountInString(hangman.NicerTypo(GameInProgress.LanguageTxt[47]))/2), hangman.NicerTypo(GameInProgress.LanguageTxt[47]))
		time.Sleep(time.Second * 2)
		MainMenu()
	}
	var Page [][]fs.FileInfo
	for i := 0; i < len(SavesFiles)/4; i++ {
		Page = append(Page, nil)
	}
	if len(SavesFiles)%4 != 0 {
		Page = append(Page, nil)
	}
	for i := 0; i < len(SavesFiles); i++ {
		Page[i/4] = append(Page[i/4], SavesFiles[i])
	}

	PrintSaves(Page, 0)
}

func PrintSaves(Page [][]fs.FileInfo, p int) {
	ClearMenu()
	for i := 232; i <= 250; i++ {
		fmt.Printf("\x1B[38;5;%dm", i)
		for j, save := range Page[p] {
			fmt.Print(MoveTo(33+3*j, w.colones/2-utf8.RuneCountInString("(n)   "+save.Name())), "(", j+1, ")   ", hangman.NicerTypo(save.Name()[:utf8.RuneCountInString(save.Name())-5]))
		}
		if p < len(Page)-1 {
			fmt.Print(MoveTo(44, w.colones/2+25), GameInProgress.LanguageTxt[42]) // NEXT
		}
		if p > 0 {
			fmt.Print(MoveTo(44, w.colones/2-25-utf8.RuneCountInString(GameInProgress.LanguageTxt[43])), GameInProgress.LanguageTxt[43]) // PREV
		}
		if p == len(Page) {
			fmt.Print(MoveTo(44, w.colones/2+25), "                    ")
		}
		if p == 0 {
			fmt.Print(MoveTo(44, w.colones/2-25-utf8.RuneCountInString(GameInProgress.LanguageTxt[43])), "                    ")
		}
		time.Sleep(time.Millisecond * 20)
	}

	var choice string
	fmt.Print(MoveTo(1, 1), ">_")
	fmt.Scanln(&choice)
	switch choice {
	case "9":
		if p < len(Page)-1 {
			PrintSaves(Page, p+1) // Next
		} else {
			PrintSaves(Page, p)
		}
	case "1":
		readJSON("Saves/" + Page[p][0].Name())
		ClearMenu()
	case "2":
		if len(Page[p]) >= 2 {
			readJSON("Saves/" + Page[p][2].Name())
			ClearMenu()
		} else {
			PrintSaves(Page, p)
		}
	case "3":
		if len(Page[p]) >= 3 {
			readJSON("Saves/" + Page[p][3].Name())
			ClearMenu()
		} else {
			PrintSaves(Page, p)
		}
	case "4":
		if len(Page[p]) >= 4 {
			readJSON("Saves/" + Page[p][4].Name())
			ClearMenu()
		} else {
			PrintSaves(Page, p)
		}
	case "8":
		if p > 0 {
			PrintSaves(Page, p-1) // Prev
		} else {
			PrintSaves(Page, p)
		}
	default:
		PrintSaves(Page, p)
	}
}
