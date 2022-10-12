package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func (GameInProgress *Game) Marshal(file string) {
	b, err := json.Marshal(GameInProgress)
	if err != nil {
		fmt.Println("error", err)
	}
	os.WriteFile(file, b, 0644)
}

func (GameInProgress *Game) readJSON(file string) *Game {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println("error", err)
	}
	json.NewDecoder(jsonFile).Decode(&GameInProgress)
	return GameInProgress
}
