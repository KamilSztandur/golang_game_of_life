package main

import (
	"fmt"
	"golang_game_of_life/config"
	"golang_game_of_life/game_engine"
	"golang_game_of_life/input_handlers"
	"os"
	"strconv"
)

func main() {
	var threadsAmount int = config.DefaultThreadsAmount

	if input_handlers.UserPassedArgument(os.Args) {
		fmt.Println("Don't forget that you can customize number of threads by passing argument to this program")
		fmt.Println("$ go build main.go 16")
		fmt.Println("Number must be a power of 4 (4, 16, 64, 256 ... etc). Default value is 4.")
	} else {
		inputAmount, err := strconv.Atoi(os.Args[1])

		if !input_handlers.AreGivenArgumentsValid(inputAmount, err) {
			fmt.Println("Invalid argument. It must be integer and power of '4' number.")
			return
		}

		if input_handlers.IsGivenThreadsAmountBiggerThanMap(inputAmount) {
			threadsAmount = config.MapSize * config.MapSize
			fmt.Printf("Entered threads amount is bigger than map which is prohibited. Reducing threads amount to %d.\n", threadsAmount)
		} else {
			threadsAmount = inputAmount
		}

	}

	game_engine.LaunchGame(threadsAmount)
}
