package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int 

func New(s string) (Matrix, error) {
	r := strings.Split(s, "\n")
	matrix := make([][]int, len(r))
	for i, row := range r {
		c := strings.Fields(row)
		if len(c) == 0 {
			return Matrix{}, errors.New("empty row")
		}
		matrix[i] = make([]int, len(c))
		if i > 0 && len(c) != len(matrix[i-1]) {
			return Matrix{}, errors.New("uneven row length")
		}
		for j, col := range c {
			n, err := strconv.Atoi(col)
			if err != nil {
				return Matrix{}, err
			}
			matrix[i][j] = n
		}
	}


	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for i := range cols {
		cols[i] = make([]int, len(m))
		for j := range cols[i] {
			cols[i][j] = m[j][i]
		}
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i := range rows {
		rows[i] = make([]int, len(m[i]))
		copy(rows[i], m[i])
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
