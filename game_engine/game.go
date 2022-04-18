package game_engine

import (
	"golang_game_of_life/communication"
	"golang_game_of_life/config"
)

var threadsAmount int
var oldMap [config.MapSize][config.MapSize]bool
var newMap [config.MapSize][config.MapSize]bool

func LaunchGame(n int) {
	threadsAmount = n
	oldMap = GenerateMap()

	runGame()
}

func runGame() {
	for true {
		communication.PrintMapState(oldMap)
		newMap = getUpdatedMap()

		oldMap = newMap
	}
}

func getUpdatedMap() [config.MapSize][config.MapSize]bool {
	return oldMap
}

func divideMapForChunks() {
	//TODO
}

func updateChunk(rowStartIndex int, rowEndIndex int, colStartIndex int, colEndIndex int) {
	//TODO
}
