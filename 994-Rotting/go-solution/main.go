package main

import (
	"fmt"
)

func hasFresh(grid [][]int) bool {
	for _, i := range grid {
		for _, j := range i {
			if j == 1 {
				return true
			}
		}
	}
	return false
}

var dirs = [][]int{
	{0, 1},  // right
	{0, -1}, // left
	{1, 0},  // down
	{-1, 0}, //up
}

// check if fresh
func check(grid [][]int, row, col int) bool {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return false
	}
	return grid[row][col] == 1
}

func orangesRotting(grid [][]int) int {
	time := 0
	// use a queue to keep track of which ones have been rotten
	var q [][]int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			}
		}
	}
	if len(q) == 0 {
		// no rotten to begin with
		if hasFresh(grid) {
			return -1
		}
		return 0
	}
	for hasFresh(grid) {
		time++
		// dequeue all, create new q
		var newQ [][]int
		for _, elem := range q {
			rrow := elem[0]
			rcol := elem[1]
			for _, dir := range dirs {
				nr := rrow + dir[0]
				nc := rcol + dir[1]
				if check(grid, nr, nc) {
					// this is rottable!
					grid[nr][nc] = 2 // rot it
					newQ = append(newQ, []int{nr, nc})
				}
			}
		}
		if len(newQ) == 0 {
			// no more rotting
			break
		}
		q = newQ
	}
	if hasFresh(grid) {
		return -1
	}
	return time
}
func main() {
	o := orangesRotting
	var t [][]int
	t = [][]int{{1, 2, 1, 1, 2, 1, 1}}
	fmt.Println(o(t))
	t = [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(o(t))
}
