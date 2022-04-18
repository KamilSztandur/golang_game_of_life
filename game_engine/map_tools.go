package game_engine

import (
	"golang_game_of_life/config"
	"math/rand"
	"time"
)

func GenerateMap() [config.MapSize][config.MapSize]bool {
	var golMap [config.MapSize][config.MapSize]bool

	rand.Seed(time.Now().UnixNano())

	for i, _ := range golMap {
		for j, _ := range golMap[i] {
			golMap[i][j] = rand.Intn(10) == 1
		}
	}

	return golMap
}

func GetCellEnvironment(golMap [config.MapSize][config.MapSize]bool, cellRowIndex int, cellColIndex int) [3][3]bool {
	var environment [3][3]bool

	environment[1][1] = golMap[cellRowIndex][cellColIndex]

	for rowIndex, row := range environment {
		for colIndex, _ := range row {
			var rowIndexDelta = rowIndex - cellRowIndex
			var colIndexDelta = colIndex - cellColIndex

			var y = cellRowIndex + rowIndexDelta
			var x = cellColIndex + colIndexDelta

			if isOutOfRange(y, x) {
				environment[rowIndex][colIndex] = false
			} else {
				environment[rowIndex][colIndex] = golMap[y][x]
			}
		}
	}

	return environment
}

func isOutOfRange(rowIndex int, colIndex int) bool {
	return colIndex < 0 || colIndex >= config.MapSize || rowIndex < 0 || rowIndex >= config.MapSize
}
