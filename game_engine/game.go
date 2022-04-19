package game_engine

import (
	"fmt"
	"golang_game_of_life/communication"
	"golang_game_of_life/config"
	"math"
)

var threadsAmount int
var chunks []ChunkIndexes

var oldMap [config.MapSize][config.MapSize]bool
var newMap [config.MapSize][config.MapSize]bool

type ChunkIndexes struct {
	rowStartIndex int
	rowEndIndex   int
	colStartIndex int
	colEndIndex   int
}

func LaunchGame(n int) {
	threadsAmount = n
	oldMap = GenerateMap()
	chunks = divideMapForChunks()

	//runGame()
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

func divideMapForChunks() []ChunkIndexes {
	chunks := make([]ChunkIndexes, threadsAmount)
	getChunksIndexesList(chunks[:])

	for _, value := range chunks {
		fmt.Println(value)
	}

	return chunks
}

func getChunksIndexesList(chunks []ChunkIndexes) {
	howManyChunksInLength := int(math.Sqrt(float64(threadsAmount)))
	chunkSize := int(float64(config.MapSize) / float64(howManyChunksInLength))

	currentChunkIndex := 0

	for i := 0; i < howManyChunksInLength; i++ {
		rowStartIndex := i * chunkSize
		rowEndIndex := (i+1)*chunkSize - 1

		for j := 0; j < howManyChunksInLength; j++ {
			colStartIndex := j * chunkSize
			colEndIndex := (j+1)*chunkSize - 1

			currentChunk := ChunkIndexes{
				rowStartIndex: rowStartIndex,
				colStartIndex: colStartIndex,
				colEndIndex:   colEndIndex,
				rowEndIndex:   rowEndIndex,
			}

			chunks[currentChunkIndex] = currentChunk

			currentChunkIndex++
		}
	}
}

func updateChunk(rowStartIndex int, rowEndIndex int, colStartIndex int, colEndIndex int) {
	//TODO
}
