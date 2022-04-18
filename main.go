package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const DefaultThreadsAmount int = 4
const MapSize int = 16

func main() {
	var threadsAmount int = DefaultThreadsAmount

	if len(os.Args) == 1 {
		fmt.Println("Don't forget that you can customize number of threads by passing argument to this program")
		fmt.Println("$ go build main.go 16")
		fmt.Println("Number must be a power of 4 (4, 16, 64, 256 ... etc). Default value is 4.")
	} else {
		inputAmount, err := strconv.Atoi(os.Args[1])

		if err != nil || !IsPowerOfFour(inputAmount) {
			fmt.Println("Invalid argument. It must be integer and power of '4' number.")
			return
		}

		threadsAmount = inputAmount
	}

	RunGame(threadsAmount)
}

func IsPowerOfFour(number int) bool {
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

func RunGame(threadsAmount int) {
	var myMap = GenerateMap()
	for _, row := range myMap {
		for _, cell := range row {
			if cell {
				fmt.Print("# ")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
	return
}

func GenerateMap() [MapSize][MapSize]bool {
	var golMap [MapSize][MapSize]bool

	rand.Seed(time.Now().UnixNano())

	for i, _ := range golMap {
		for j, _ := range golMap[i] {
			golMap[i][j] = rand.Intn(10) == 1
		}
	}

	return golMap
}

func UpdateChunk(chunk [][]bool) [][]bool {
	updatedChunk := chunk

	for rowIndex, row := range chunk {
		for colIndex, cell := range row {
			var cellEnvironment = GetCellEnvironment(chunk, rowIndex, colIndex)

			if cell {
				updatedChunk[rowIndex][colIndex] = ShouldCellSurvive(cellEnvironment)
			} else {
				updatedChunk[rowIndex][colIndex] = ShouldCellComeToLife(cellEnvironment)
			}
		}
	}

	return updatedChunk
}

func GetCellEnvironment(golMap [MapSize][MapSize]bool, cellRowIndex int, cellColIndex int) [3][3]bool {
	var environment [3][3]bool

	environment[1][1] = golMap[cellRowIndex][cellColIndex]

	for rowIndex, row := range environment {
		for colIndex, _ := range row {
			var rowIndexDelta = rowIndex - cellRowIndex
			var colIndexDelta = colIndex - cellColIndex

			var y = cellRowIndex + rowIndexDelta
			var x = cellColIndex + colIndexDelta

			if IsOutOfRange(y, x) {
				environment[rowIndex][colIndex] = false
			} else {
				environment[rowIndex][colIndex] = golMap[y][x]
			}
		}
	}

	return environment
}

func IsOutOfRange(rowIndex int, colIndex int) bool {
	if colIndex < 0 || colIndex >= MapSize {
		return true
	} else if rowIndex < 0 || rowIndex >= MapSize {
		return true
	} else {
		return false
	}
}

func ShouldCellSurvive(fragment [3][3]bool) bool {
	var aliveNeighbors = 0

	if !fragment[1][1] {
		return false
	}

	for rowIndex, row := range fragment {
		for colIndex, cell := range row {
			if colIndex == 1 && rowIndex == 1 {
				continue
			}

			if cell {
				aliveNeighbors++
			}
		}
	}

	return aliveNeighbors == 2 || aliveNeighbors == 3
}

func ShouldCellComeToLife(fragment [3][3]bool) bool {
	var aliveNeighbors = 0

	if fragment[1][1] {
		return false
	}

	for rowIndex, row := range fragment {
		for colIndex, cell := range row {
			if colIndex == 1 && rowIndex == 1 {
				continue
			}

			if cell {
				aliveNeighbors++
			}
		}
	}

	return aliveNeighbors == 3
}
