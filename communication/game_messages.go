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

const cellCharacter = string(rune(128))

func PrintHelpMessage() {
	fmt.Println("Don't forget that you can customize number of threads by passing argument to this program")
	fmt.Println("$ go build main.go 16")
	fmt.Printf("Number must be a power of 4 (4, 16, 64, 256 ... etc). Default value is %d.", config.DefaultThreadsAmount)
}

func PrintInvalidArgumentsMessage() {
	fmt.Println("Invalid argument. It must be integer and power of '4' number.")
}

func PrintThreadsAmountIsBiggerThanMapMessage(newThreadsAmount int) {
	fmt.Printf("Entered threads amount is bigger than map which is prohibited. Reducing threads amount to %d.\n", newThreadsAmount)
}

func PrintMapState(currentMap [config.MapSize][config.MapSize]bool) {
	printMapToScreen(currentMap)
	waitFewMoments()
	clearScreen()
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

func clearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}

func printMapToScreen(currentMap [config.MapSize][config.MapSize]bool) {
	for _, row := range currentMap {
		for _, cell := range row {
			if cell {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func waitFewMoments() {
	time.Sleep(1 * time.Second)
}
