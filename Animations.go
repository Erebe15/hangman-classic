package main

import (
	"fmt"
	"math"
	"time"
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

		case 2:

		case 3:

		case 4:
			SpawnCorde(w.colones*85/100, w.ligns*15/100, 20)
		case 5:
		case 6:
		case 7:
		case 8:
		case 9:
		case 10:
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

func RemoveLine(xa, ya, xb, yb int) {
	if xa > xb {
		xa, ya, xb, yb = xb, yb, xa, ya
	}
	dx := xb - xa
	dy := yb - ya
	for i := xa; i < xb; i++ {
		j := ya + dy*(i-xa)/dx
		fmt.Print(MoveTo(j, i), " ")
	}
	if ya > yb {
		xa, ya, xb, yb = xb, yb, xa, ya
	}
	for i := ya; i < yb; i++ {
		j := xa + dx*(i-ya)/dy
		fmt.Print(MoveTo(i, j), " ")
	}
}

func lerp(A, B, t float64) float64 {
	return A + (B-A)*t
}

func Elastic(x float64) float64 {
	return math.Sin(-13*(x+1)*math.Pi/2)*math.Pow(2, -10*x) + 1
}

func SpawnCorde(x, y, l int) {
	for i := 1; i <= 200; i++ {
		t0 := float64(i-1) / 100
		t := float64(i) / 100
		RemoveNode(x, y+int(lerp(0, float64(w.ligns*l/100)+0.2, Elastic(t0))))
		RemoveLine(x, y, x, y+int(lerp(0, float64(w.ligns*l/100)+0.2, Elastic(t0))))
		DrawNode(x, y+int(lerp(0, float64(w.ligns*l/100)+0.2, Elastic(t))))
		DrawLine(x, y, x, y+int(lerp(0, float64(w.ligns*l/100)+0.2, Elastic(t))), "█")
		time.Sleep(time.Millisecond)
	}
}

func RemoveNode(x, y int) {
	fmt.Print(MoveTo(y, x), " ")
	fmt.Print("\x1B[B\x1B[3D     ")
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
	fmt.Print("\x1B[B\x1B[7D█▀   ▀█")
	fmt.Print("\x1B[B\x1B[7D█▄   ▄█")
	fmt.Print("\x1B[B\x1B[7D █▄▄▄█ ")
}
