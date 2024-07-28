package spiralmatrix

func SpiralMatrix(size int) [][]int {
	matrix := make([][]int, size)
	if size == 0 {
		return matrix
	}
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	matrix[0][0] = 1
	directions := []string{"r", "d", "l", "u"}

	var row, col int
	for i := 0; i < 2*size-1; i++ {
		direction := directions[i%4]
		matrix, row, col = fillInDirection(matrix, direction, row, col)
	}

	return matrix
}

func fillInDirection(matrix [][]int, direction string, row, col int) ([][]int, int, int) {
	num := matrix[row][col] + 1
	finalRow, finalCol := row, col

	var step func(int, int) (int, int)

	switch direction {
	case "r":
		step = func(r, c int) (int, int) { return r, c + 1 }
	case "l":
		step = func(r, c int) (int, int) { return r, c - 1 }
	case "u":
		step = func(r, c int) (int, int) { return r - 1, c }
	case "d":
		step = func(r, c int) (int, int) { return r + 1, c }
	}

	r, c := step(row, col)
	for r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[0]) && matrix[r][c] == 0 {
		matrix[r][c] = num
		num++
		finalRow, finalCol = r, c
		r, c = step(r, c)
	}

	return matrix, finalRow, finalCol
}
