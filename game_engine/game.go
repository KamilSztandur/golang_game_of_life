package game_engine

import (
	"golang_game_of_life/communication"
	"golang_game_of_life/config"
	"sync"
)

type GameState struct {
	chunks     []ChunkIndexes
	currentMap [config.MapSize][config.MapSize]bool
}

func LaunchGame(threadsAmount int) {
	state := GameState{
		chunks:     DivideMapForChunks(threadsAmount),
		currentMap: GenerateMap(),
	}

	runGame(&state)
}

func runGame(state *GameState) {
	for {
		communication.PrintMapState(state.currentMap)
		newMap := getUpdatedMap(state)

		state.currentMap = newMap
	}
}

func getUpdatedMap(state *GameState) [config.MapSize][config.MapSize]bool {
	var updatedMap [config.MapSize][config.MapSize]bool
	var waitGroup sync.WaitGroup

	var threadsAmount = len(state.chunks)
	for i := 0; i < threadsAmount; i++ {
		waitGroup.Add(1)

		go func(index int) {
			defer waitGroup.Done()
			updateChunk(index, state, &updatedMap)
		}(i)
	}

	waitGroup.Wait()

	return updatedMap
}

func updateChunk(chunkIndex int, state *GameState, newMap *[config.MapSize][config.MapSize]bool) {
	rowStartIndex := state.chunks[chunkIndex].rowStartIndex
	rowEndIndex := state.chunks[chunkIndex].rowEndIndex
	colStartIndex := state.chunks[chunkIndex].colStartIndex
	colEndIndex := state.chunks[chunkIndex].colEndIndex

	for rowIndex := rowStartIndex; rowIndex <= rowEndIndex; rowIndex++ {
		for colIndex := colStartIndex; colIndex <= colEndIndex; colIndex++ {
			cellEnvironment := GetCellEnvironment(&state.currentMap, rowIndex, colIndex)

			if state.currentMap[rowIndex][colIndex] {
				newMap[rowIndex][colIndex] = ShouldCellSurvive(cellEnvironment)
			} else {
				newMap[rowIndex][colIndex] = ShouldCellComeToLife(cellEnvironment)
			}
		}
	}
}
