package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Marshal(save Game) {
	b, err := json.Marshal(save)
	if err != nil {
		fmt.Println("error", err)
	}
	os.WriteFile("save.json", b, 0644)
	//os.WriteFile("save.txt", b, 0644)

}

func PrintSave(stry string) {

	content, err := os.ReadFile("save.json")
	if err != nil {
		fmt.Println("error", err)

	}
	var save Game
	err = json.Unmarshal(content, &save)
	if err != nil {
		fmt.Println("error", err)
	}
}
