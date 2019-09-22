package main

import (
	"fmt"
)

func check(maze [][]int, row, col int) bool {
	maxRow := len(maze)
	maxCol := len(maze[0])
	if row < 0 || row >= maxRow || col < 0 || col >= maxCol {
		return false
	}
	return maze[row][col] == 0
}

var pos = [][]int{
	{0, 1},  // right
	{0, -1}, // left
	{1, 0},  // down
	{-1, 0}, //up
}

// Depth first search
func walk(maze [][]int, row, col int, solution *[][]bool) bool {
	maxRow := len(maze) - 1
	maxCol := len(maze[0]) - 1

	if row == maxRow && col == maxCol {
		fmt.Println("Solved!")
		(*solution)[row][col] = true
		return true //done
	}
	for _, dir := range pos {
		nr := row + dir[0]
		nc := col + dir[1]
		if check(maze, nr, nc) {
			fmt.Printf("Walking: %d, %d\n", nr, nc)
			// ok! we can walk there
			maze[row][col] = 2 // mark our walked place
			(*solution)[row][col] = true
			if walk(maze, nr, nc, solution) {
				// woah That's our path!
				return true
			}
			// nope, set back
			maze[row][col] = 0
			(*solution)[row][col] = false
		}
	}
	return false
}

func solveMaze(maze [][]int) [][]bool {
	maxRow := len(maze)
	maxCol := len(maze[0])
	solution := make([][]bool, maxRow)
	for i := 0; i < maxRow; i++ {
		solution[i] = make([]bool, maxCol)
	}
	walk(maze, 0, 0, &solution)
	return solution
}

func main() {
	maze := [][]int{
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
		{1, 1, 0, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 1},
		{1, 1, 0, 0, 0},
	}
	ans := solveMaze(maze)
	for _, e := range ans {
		fmt.Println(e)
	}
}
