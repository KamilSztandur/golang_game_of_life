package game_engine

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
