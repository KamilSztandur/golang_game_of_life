package input_handlers

import (
	"golang_game_of_life/config"
	"os"
)

func UserPassedArgument() bool {
	return len(os.Args) == 1
}

func AreGivenArgumentsValid(parsedNumber int, err error) bool {
	if err != nil || !isPowerOfFour(parsedNumber) {
		return false
	}

	return true
}

func IsGivenThreadsAmountBiggerThanMap(threadsAmount int) bool {
	return threadsAmount > config.MapSize^2
}

func isPowerOfFour(number int) bool {
	if number == 0 {
		return false
	}

	for number != 1 {
		if number%4 != 0 {
			return false
		}

		number = number / 4
	}

	return true
}
