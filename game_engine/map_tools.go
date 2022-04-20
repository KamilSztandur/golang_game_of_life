package game_engine

import (
	"golang_game_of_life/config"
	"math"
	"math/rand"
	"time"
)

type ChunkIndexes struct {
	rowStartIndex int
	rowEndIndex   int
	colStartIndex int
	colEndIndex   int
}

func GenerateMap() [config.MapSize][config.MapSize]bool {
	var golMap [config.MapSize][config.MapSize]bool

	rand.Seed(time.Now().UnixNano())

	for rowIndex, row := range golMap {
		for colIndex := range row {
			golMap[rowIndex][colIndex] = rand.Intn(2) != 1
		}
	}

	return golMap
}

func GetCellEnvironment(golMap *[config.MapSize][config.MapSize]bool, cellRowIndex int, cellColIndex int) [3][3]bool {
	var environment [3][3]bool

	for rowIndex, row := range environment {
		for colIndex := range row {

			y := cellRowIndex - (1 - rowIndex)
			x := cellColIndex - (1 - colIndex)

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

func DivideMapForChunks(threadsAmount int) []ChunkIndexes {
	chunks := make([]ChunkIndexes, threadsAmount)
	calculateChunksIndexesList(chunks[:])

	return chunks
}

func calculateChunksIndexesList(chunks []ChunkIndexes) {
	threadsAmount := len(chunks)

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
