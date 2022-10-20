package main

import (
	"fmt"
	"os"
	"strings"
)

func PrintJose() {
	data, err := os.ReadFile("FilesAndLists/hangman.txt")
	if err != nil {
		panic(err)
	}
	var JoseFloor [][]string
	States := strings.Split(string(data), "\n\n")
	for _, i := range States {
		JoseFloor = append(JoseFloor, strings.Split(i, "\n"))
	}
	for j, l := range JoseFloor[GameInProgress.Tries] {
		fmt.Print(MoveTo(w.ligns/2-4+j, w.colones*75/100), l)
	}
}

func DrawHangmanState() {
	if w.ligns >= 25 && w.colones > 65 {
		drawFancy()
	} else {
		PrintJose()
	}
}

func drawFancy() {
	switch GameInProgress.Tries {
	case 10:
		DrawLine(w.colones*85/100, w.ligns*55/100, w.colones*85/100+(w.ligns/10+2), w.ligns*55/100+(w.ligns/10+2), "█") // Right leg
		fallthrough
	case 9:
		DrawLine(w.colones*85/100, w.ligns*55/100, w.colones*85/100-(w.ligns/10+1), w.ligns*55/100+(w.ligns/10+1), "█") // left leg
		fallthrough
	case 8:
		DrawLine(w.colones*85/100, w.ligns*15/100+w.ligns*20/100+5, w.colones*85/100+(w.ligns/10+2), w.ligns*15/100+w.ligns*20/100+5+(w.ligns/10+2), "█") // Right arm
		fallthrough
	case 7:
		DrawLine(w.colones*85/100, w.ligns*15/100+w.ligns*20/100+5, w.colones*85/100-(w.ligns/10+1), w.ligns*15/100+w.ligns*20/100+5+(w.ligns/10+1), "█") // Left arm
		fallthrough
	case 6:
		DrawLine(w.colones*85/100, w.ligns*15/100+w.ligns*20/100+5, w.colones*85/100, w.ligns*55/100, "█") // Body
		fallthrough
	case 5:
		if GameInProgress.Tries == 5 {
			RemoveNode(w.colones*85/100, w.ligns*15/100+w.ligns*20/100) // Head
		}
		if GameInProgress.Tries >= 10 {
			DrawDeadHead(w.colones*85/100, w.ligns*15/100+w.ligns*20/100)
		} else if GameInProgress.Tries == 9 {
			DrawFreakOutHead(w.colones*85/100, w.ligns*15/100+w.ligns*20/100)
		} else {
			DrawHead(w.colones*85/100, w.ligns*15/100+w.ligns*20/100)
		}
		fallthrough
	case 4:
		DrawLine(w.colones*85/100, w.ligns*15/100, w.colones*85/100, w.ligns*15/100+w.ligns*20/100, "█") // Rope
		if GameInProgress.Tries == 4 {
			DrawNode(w.colones*85/100, w.ligns*15/100+w.ligns*20/100)
		}
		fallthrough
	case 3:
		DrawLine(w.colones*71/100, w.ligns*15/100, w.colones*87/100, w.ligns*15/100, "█") // Transverse
		DrawLine(w.colones*73/100, w.ligns*25/100, w.colones*78/100, w.ligns*15/100, "█")
		fallthrough
	case 2:
		DrawLine(w.colones*73/100, w.ligns*90/100, w.colones*73/100, w.ligns*15/100, "█") // Pole
		DrawLine(w.colones*73/100-1, w.ligns*90/100, w.colones*73/100-1, w.ligns*15/100+1, "▒")
		DrawLine(w.colones*73/100-2, w.ligns*90/100, w.colones*73/100-2, w.ligns*15/100, "█")
		fallthrough
	case 1:
		DrawLine(w.colones*70/100, w.ligns*90/100, w.colones*95/100, w.ligns*90/100, "█") // Ground
		DrawLine(w.colones*70/100, w.ligns*90/100+1, w.colones*95/100, w.ligns*90/100+1, "█")
		DrawLine(w.colones*70/100+1, w.ligns*90/100+1, w.colones*95/100-1, w.ligns*90/100+1, "▒")
	}
}
