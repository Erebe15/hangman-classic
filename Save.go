package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func WriteFile(saves Game) {
	message := fmt.Sprintf("%d", saves)
	err := ioutil.WriteFile("saves.txt", []byte(message), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
