package main

import (
	"fmt"
	"math"
	"time"
	"unicode/utf8"
	hangman "ytrack.learn.ynov.com/git/pthomas/hangman-classic/Functions"
)

func WrongAnswerEffect() {
	for i := 0; i <= 6; i++ {
		if i%2 == 1 {
			fmt.Print("\x1B[31m")
		} else {
			fmt.Print("\x1B[0m")
		}
		CreateWindows()
		time.Sleep(time.Second / 10)
	}
	fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[0m")
}

func GoodAnswerEffect() {
	fmt.Print("\x1B[32m")
	CreateWindows()
	fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[0m")
	time.Sleep(time.Second / 3)
	CreateWindows()
}

func HangmanAnimation() {
	if w.ligns >= 25 && w.colones > 65 {
		switch GameInProgress.Tries {
		case 1:
			for i := w.ligns * 25 / 100; i <= w.ligns*90/100; i++ {
				DrawLine(w.colones*82/100-i/3, w.ligns*90/100, w.colones*82/100+i/3, w.ligns*90/100, "▂")
				DrawLine(w.colones*70/100, i, w.colones*95/100, i, "█") // Ground
				DrawLine(w.colones*70/100, i+1, w.colones*95/100, i+1, "█")
				DrawLine(w.colones*70/100+1, i+1, w.colones*95/100-1, i+1, "▒")
				time.Sleep(time.Millisecond * time.Duration(100/i))
				DrawLine(w.colones*70/100, i, w.colones*95/100, i, " ") // Ground
				DrawLine(w.colones*70/100, i+1, w.colones*95/100, i+1, " ")
			}
			DrawLine(w.colones*70/100, w.ligns*90/100, w.colones*95/100, w.ligns*90/100, "█") // Ground
			DrawLine(w.colones*70/100, w.ligns*90/100+1, w.colones*95/100, w.ligns*90/100+1, "█")
			DrawLine(w.colones*70/100+1, w.ligns*90/100+1, w.colones*95/100-1, w.ligns*90/100+1, "▒")
		case 2:
			for i := 0; i <= w.ligns*75/100; i++ {
				DrawLine(w.colones*73/100, w.ligns*90/100, w.colones*73/100, w.ligns*90/100-i, "█") // Pole
				DrawLine(w.colones*73/100-1, w.ligns*90/100, w.colones*73/100-1, w.ligns*90/100-i+1, "▒")
				DrawLine(w.colones*73/100-2, w.ligns*90/100, w.colones*73/100-2, w.ligns*90/100-i, "█")
				time.Sleep(time.Millisecond * 5)
			}
		case 3:
			for i := 314 / 4; i >= 0; i-- {
				DrawLine(w.colones*73/100, w.ligns*90/100, w.colones*73/100, w.ligns*15/100, "█")
				DrawLine(w.colones*73/100-1, w.ligns*90/100, w.colones*73/100-1, w.ligns*15/100+1, "▒")
				DrawLine(w.colones*73/100-2, w.ligns*90/100, w.colones*73/100-2, w.ligns*15/100, "█")
				DrawLine(w.colones*71/100, w.ligns*15/100, (w.colones*71/100)+int(math.Cos(float64(i)/50)*float64(w.colones*16/100)), w.ligns*15/100+int(math.Sin(float64(i)/50)*float64(w.colones*8/100)), "█")
				time.Sleep(time.Millisecond)
				DrawLine(w.colones*71/100, w.ligns*15/100, (w.colones*71/100)+int(math.Cos(float64(i)/50)*float64(w.colones*16/100)), w.ligns*15/100+int(math.Sin(float64(i)/50)*float64(w.colones*8/100)), " ")
			}
			DrawLine(w.colones*71/100, w.ligns*15/100, w.colones*87/100, w.ligns*15/100, "█") // Transverse
			time.Sleep(time.Millisecond * 100)
			DrawLine(w.colones*73/100, w.ligns*25/100, w.colones*78/100, w.ligns*15/100, "█")
		case 4:
			SpawnCorde(w.colones*85/100, w.ligns*15/100, 20)
		case 5:
			RemoveNode(w.colones*85/100, w.ligns*15/100+w.ligns*20/100)
			DrawHead(w.colones*85/100, w.ligns*15/100+w.ligns*20/100)
		case 6:
			for i := w.ligns*15/100 + w.ligns*20/100 + 5; i <= w.ligns*55/100; i++ {
				DrawLine(w.colones*85/100, w.ligns*15/100+w.ligns*20/100+5, w.colones*85/100, i, "█") // Body
				time.Sleep(time.Millisecond * 100)
			}
		case 7:
			for i := 0; i <= (w.ligns/10 + 1); i++ {
				DrawLine(w.colones*85/100, w.ligns*15/100+w.ligns*20/100+5, w.colones*85/100-i, w.ligns*15/100+w.ligns*20/100+5+i, "█") // Left arm
				time.Sleep(time.Millisecond * 100)
			}
		case 8:
			for i := 0; i <= (w.ligns/10 + 2); i++ {
				DrawLine(w.colones*85/100, w.ligns*15/100+w.ligns*20/100+5, w.colones*85/100+i, w.ligns*15/100+w.ligns*20/100+5+i, "█")
				time.Sleep(time.Millisecond * 100)
			}
		case 9:
			for i := 0; i <= (w.ligns/10 + 1); i++ {
				DrawLine(w.colones*85/100, w.ligns*55/100, w.colones*85/100-i, w.ligns*55/100+i, "█") // left leg
				time.Sleep(time.Millisecond * 100)
			}
		case 10:
			for i := 0; i <= (w.ligns/10 + 2); i++ {
				DrawLine(w.colones*85/100, w.ligns*55/100, w.colones*85/100+i, w.ligns*55/100+i, "█") // Right leg
				time.Sleep(time.Millisecond * 100)
			}
		}
	}
}

func DrawLine(xa, ya, xb, yb int, c string) {
	if xa > xb {
		xa, ya, xb, yb = xb, yb, xa, ya
	}
	dx := xb - xa
	dy := yb - ya
	for i := xa; i < xb; i++ {
		j := ya + dy*(i-xa)/dx
		fmt.Print(MoveTo(j, i), c)
	}
	if ya > yb {
		xa, ya, xb, yb = xb, yb, xa, ya
	}
	for i := ya; i < yb; i++ {
		j := xa + dx*(i-ya)/dy
		fmt.Print(MoveTo(i, j), c)
	}
}

func lerp(A, B, t float64) float64 {
	return A + (B-A)*t
}

func Elastic(x float64) float64 {
	return math.Sin(-13*(x+1)*math.Pi/2)*math.Pow(2, -10*x) + 1
}

func SpawnCorde(x, y, l int) {
	for i := 1; i <= 150; i++ {
		t0 := float64(i-1) / 100
		t := float64(i) / 100
		RemoveNode(x, y+int(lerp(0, float64(w.ligns*l/100)+0.2, Elastic(t0))))
		DrawLine(x, y, x, y+int(lerp(0, float64(w.ligns*l/100)+0.2, Elastic(t0))), " ")
		DrawNode(x, y+int(lerp(0, float64(w.ligns*l/100)+0.2, Elastic(t))))
		DrawLine(x, y, x, y+int(lerp(0, float64(w.ligns*l/100)+0.2, Elastic(t))), "█")
		time.Sleep(time.Millisecond)
	}
}

func RemoveNode(x, y int) {
	fmt.Print(MoveTo(y, x), "")
	fmt.Print("\x1B[B\x1B[2D     ")
	fmt.Print("\x1B[B\x1B[5D     ")
	fmt.Print("\x1B[B\x1B[5D     ")
	fmt.Print("\x1B[B\x1B[5D     ")
	fmt.Print("\x1B[B\x1B[5D    ")
}

func DrawNode(x, y int) {
	fmt.Print(MoveTo(y, x), "█")
	fmt.Print("\x1B[B\x1B[3D █ █ ")
	fmt.Print("\x1B[B\x1B[5D█   █")
	fmt.Print("\x1B[B\x1B[5D█   █")
	fmt.Print("\x1B[B\x1B[5D█   █")
	fmt.Print("\x1B[B\x1B[5D █▄█")
}

func DrawHead(x, y int) {
	fmt.Print(MoveTo(y, x), "\x1B[3D   █   ")
	fmt.Print("\x1B[B\x1B[7D █▀▀▀█ ")
	fmt.Print("\x1B[B\x1B[7D█▀● ●▀█")
	fmt.Print("\x1B[B\x1B[7D█▄ ⁀ ▄█")
	fmt.Print("\x1B[B\x1B[7D █▄▄▄█ ")
}

func DrawDeadHead(x, y int) {
	fmt.Print(MoveTo(y, x), "\x1B[3D   █   ")
	fmt.Print("\x1B[B\x1B[7D █▀▀▀█ ")
	fmt.Print("\x1B[B\x1B[7D█▀X X▀█")
	fmt.Print("\x1B[B\x1B[7D█▄   ▄█")
	fmt.Print("\x1B[B\x1B[7D █▄▄▄█ ")
}

func DrawFreakOutHead(x, y int) {
	fmt.Print(MoveTo(y, x), "\x1B[3D   █   ")
	fmt.Print("\x1B[B\x1B[7D █▀▀▀█ ")
	fmt.Print("\x1B[B\x1B[7D█▀. .▀█")
	fmt.Print("\x1B[B\x1B[7D█▄   ▄█")
	fmt.Print("\x1B[B\x1B[7D █▄▄▄█ ")
}

func WelcomeArrive() {
	w.ligns, w.colones = hangman.Size()
	hangman.Clear()
	time.Sleep(time.Second)
	for i := 232; i <= 250; i++ {
		fmt.Printf("\x1B[38;5;%dm\x1B[?25l", i)
		PrintAscii(3, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[0])*12)/2, GameInProgress.LanguageTxt[0]) // WELCOME
		fmt.Printf("\x1B[38;5;%dm", i)
		PrintAscii(12, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[1])*12)/2, GameInProgress.LanguageTxt[1]) // TO
		fmt.Printf("\x1B[38;5;%dm", i)
		PrintAscii(21, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[4])*12)/2, GameInProgress.LanguageTxt[4]) // HANGMAN
		time.Sleep(time.Millisecond * 80)
	}
	PrintAscii(3, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[0])*12)/2, GameInProgress.LanguageTxt[0])
	PrintAscii(12, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[1])*12)/2, GameInProgress.LanguageTxt[1])
	PrintAscii(21, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[4])*12)/2, GameInProgress.LanguageTxt[4])
	SpawnCorde(20, 1, 60)
}

func WelcomeFadeAway() {
	for i := 250; i >= 232; i-- {
		fmt.Printf("\x1B[38;5;%dm", i)
		if GameInProgress.Status < 10 {
			fmt.Print(MoveTo(w.ligns-1, w.colones-utf8.RuneCountInString("(10) "+hangman.NicerTypo(GameInProgress.LanguageTxt[45]))), "(10) "+hangman.NicerTypo(GameInProgress.LanguageTxt[45]))
			fmt.Print(MoveTo(36, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[28])/2), GameInProgress.LanguageTxt[28])
			fmt.Print(MoveTo(39, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[29])/2), GameInProgress.LanguageTxt[29])
			fmt.Print(MoveTo(42, w.colones/2-utf8.RuneCountInString(GameInProgress.LanguageTxt[30])/2), GameInProgress.LanguageTxt[30])
		}
		PrintAscii(3, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[0])*12)/2, GameInProgress.LanguageTxt[0])
		fmt.Printf("\x1B[38;5;%dm", i)
		PrintAscii(12, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[1])*12)/2, GameInProgress.LanguageTxt[1])
		fmt.Printf("\x1B[38;5;%dm", i)
		PrintAscii(21, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[4])*12)/2, GameInProgress.LanguageTxt[4])
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Printf("\x1B[38;5;0m")
	PrintAscii(3, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[0])*12)/2, GameInProgress.LanguageTxt[0])
	fmt.Printf("\x1B[38;5;0m")
	PrintAscii(12, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[1])*12)/2, GameInProgress.LanguageTxt[1])
	fmt.Printf("\x1B[38;5;0m")
	PrintAscii(21, w.colones/2-(utf8.RuneCountInString(GameInProgress.LanguageTxt[4])*12)/2, GameInProgress.LanguageTxt[4])
	for i := 250; i >= 232; i-- {
		fmt.Printf("\x1B[38;5;%dm", i)
		DrawNode(20, w.ligns*60/100+1)
		DrawLine(20, 1, 20, w.ligns*60/100+1, "█")
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Printf("\x1B[0m")
}
