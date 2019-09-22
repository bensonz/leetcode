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

func isSolved(maze [][]int, row, col int) bool {
	maxRow := len(maze) - 1
	maxCol := len(maze[0]) - 1
	if row == maxRow && col == maxCol {
		return true
	}
	return false
}

var dirs = [][]int{
	{0, 1},  // right
	{0, -1}, // left
	{1, 0},  // down
	{-1, 0}, //up
}

func walk(maze [][]int, row, col int, solution *[][]bool) bool {
	q := [][]int{{row, col}}     // initialize our queue
	(*solution)[row][col] = true //mark first
	for {                        // while queue not empty
		if len(q) == 0 {
			break
		}
		// dequeue
		curRow := q[0][0]
		curCol := q[0][1]
		q = q[1:]
		(*solution)[curRow][curCol] = true
		maze[curRow][curCol] = 2
		// check condition
		if isSolved(maze, curRow, curCol) {
			return true
		}
		// add to action queue
		deadend := 0
		for _, dir := range dirs {
			nr := curRow + dir[0]
			nc := curCol + dir[1]
			// fmt.Printf("Checking: %d,%d\n", nr, nc)
			if check(maze, nr, nc) {
				q = append(q, []int{nr, nc})
				// no need to unmark on BFS
			} else {
				deadend++
				// fmt.Printf(" Deadend: %d,%d | %d \n", nr, nc, deadend)
			}
		}
		if deadend == 4 {
			fmt.Printf("Unmarking %d, %d \n", curRow, curCol)
			(*solution)[curRow][curCol] = false
		}
	}
	(*solution)[row][col] = false //unmark first
	return false
}

func solveMaze(maze [][]int) [][]bool {
	maxRow := len(maze)
	maxCol := len(maze[0])
	solution := make([][]bool, maxRow)
	for i := 0; i < maxRow; i++ {
		solution[i] = make([]bool, maxCol)
	}
	if walk(maze, 0, 0, &solution) {
		fmt.Println("Has solution.")
	}
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
