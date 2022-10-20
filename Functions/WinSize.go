package hangman

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Size() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 30, 150
	}
	list := strings.Split(string(out[:len(out)-1]), " ")
	a, err := strconv.Atoi(list[0])
	b, err := strconv.Atoi(list[1])
	return a - 1, b - 1
}
