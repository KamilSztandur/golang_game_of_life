package communication

import (
	"fmt"
	"golang_game_of_life/config"
	"os"
	"os/exec"
	"runtime"
	"time"
)

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

func clearScreen() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			return
		}
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			return
		}

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
