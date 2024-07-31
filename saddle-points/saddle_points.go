package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix struct {
	rows [][]int
}
type Pair struct {
	row int
	column int
}

func New(s string) (*Matrix, error) {
	if s == "" {
		return &Matrix{}, nil
	}
	var rows [][]int
	for _, r := range strings.Split(s, "\n") {
		var row []int
		for _, c := range strings.Split(r, " ") {
			num, err := strconv.Atoi(c)
			if err != nil {
				return nil, fmt.Errorf("failed to parse matrix: %v", err)
			}
			row = append(row, num)
		}
		if len(row) > 0 {
			rows = append(rows, row)
		} else {
			return nil, fmt.Errorf("failed to parse matrix: empty row")
		}
	}
	return &Matrix{rows}, nil
}

func (m *Matrix) Saddle() []Pair {
	var pairs []Pair
	for i, row := range m.rows {
		for j, element := range row {
			if isHighestInRow(row,element) && isSmallestInColumn(m.rows,j,element) {
				pairs = append(pairs,Pair{i+1,j+1})
			}
		}
	}
	return pairs
}

func isHighestInRow(row []int,element int) bool {
	for _, e := range row {
		if e > element {
			return false
		}
	}
	return true
}

func isSmallestInColumn(matrix [][]int, column int,element int) bool {
	for _, row := range matrix {
		if row[column] < element {
			return false
		}
	}
	return true
}
