package minesweeper

// Annotate returns an annotated board
func Annotate(board []string) []string {
	rows := len(board)
	if rows == 0 {
		return []string{}
	}
	cols := len(board[0])
	if cols == 0 {
		return board
	}

	// Convert input board to a mutable 2D slice of runes
	runeBoard := make([][]rune, rows)
	for i := range board {
		runeBoard[i] = []rune(board[i])
	}

	// Directions for adjacent cells (8 directions)
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},          {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Annotate the board
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if runeBoard[i][j] == ' ' {
				count := 0
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < rows && nj >= 0 && nj < cols && runeBoard[ni][nj] == '*' {
						count++
					}
				}
				if count > 0 {
					runeBoard[i][j] = rune('0' + count)
				}
			}
		}
	}

	// Convert the mutable rune board back to a slice of strings
	result := make([]string, rows)
	for i := range runeBoard {
		result[i] = string(runeBoard[i])
	}

	return result
}
