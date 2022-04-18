package input_handlers

import (
	"os"
)

func UserPassedArgument(args []string) bool {
	return len(os.Args) == 1
}

func AreGivenArgumentsValid(parsedNumber int, err error) bool {
	if err != nil || !isPowerOfFour(parsedNumber) {
		return false
	}

	return true
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
