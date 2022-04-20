package main

import (
	"golang_game_of_life/communication"
	"golang_game_of_life/config"
	"golang_game_of_life/game_engine"
	"golang_game_of_life/input_handlers"
	"os"
	"strconv"
)

func main() {
	var threadsAmount int = config.DefaultThreadsAmount

	if input_handlers.UserPassedArgument(os.Args) {
		communication.PrintHelpMessage()
	} else {
		inputAmount, err := strconv.Atoi(os.Args[1])

		if !input_handlers.AreGivenArgumentsValid(inputAmount, err) {
			communication.PrintInvalidArgumentsMessage()
			return
		}

		if input_handlers.IsGivenThreadsAmountBiggerThanMap(inputAmount) {
			threadsAmount = config.MapSize * config.MapSize
			communication.PrintThreadsAmountIsBiggerThanMapMessage(threadsAmount)
		} else {
			threadsAmount = inputAmount
		}

	}

	game_engine.LaunchGame(threadsAmount)
}
