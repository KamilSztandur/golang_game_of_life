package communication

import (
	"fmt"
	"golang_game_of_life/config"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func()

func PrintMapState(currentMap [config.MapSize][config.MapSize]bool) {
	for _, row := range currentMap {
		for _, cell := range row {
			if cell {
				fmt.Print("# ")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	time.Sleep(1 * time.Second)
	callClear()
}

func init() {
	clear = make(map[string]func())

	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			return
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			return
		}
	}
}

func callClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}