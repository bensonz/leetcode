package main

import (
	"fmt"
)

// Assume n >= 1
func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	// Initialize
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	var dirs = [][]int{
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
		{-1, 0}, //up
	}
	// walk through it
	curRow := 0
	curCol := 0
	curDir := 0
	for count := 1; count < n*n+1; count++ {
		result[curRow][curCol] = count
		shouldSwerve := false
		// check if next positions is ok
		switch curDir {
		case 0:
			nc := curCol + dirs[curDir][1]
			if nc >= n || result[curRow][nc] != 0 {
				shouldSwerve = true
			}
		case 1:
			nr := curRow + dirs[curDir][0]
			if nr >= n || result[nr][curCol] != 0 {
				shouldSwerve = true
			}
		case 2:
			nc := curCol + dirs[curDir][1]
			if nc < 0 || result[curRow][nc] != 0 {
				shouldSwerve = true
			}
		case 3:
			nr := curRow + dirs[curDir][0]
			if nr < 0 || result[nr][curCol] != 0 {
				shouldSwerve = true
			}
		}

		if shouldSwerve {
			curDir = (curDir + 1) % 4
		}
		curRow += dirs[curDir][0]
		curCol += dirs[curDir][1]
	}

	return result
}

func main() {
	// fmt.Println(generateMatrix(1))
	// fmt.Println(generateMatrix(2))
	// fmt.Println(generateMatrix(3))

	fmt.Println(generateMatrix(5))

}
