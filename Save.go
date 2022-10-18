package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	if err != nil {
		fmt.Println("error === ", err)
	}
	json.NewDecoder(jsonFile).Decode(&GameInProgress)
	err = os.Remove("Saves")
}
