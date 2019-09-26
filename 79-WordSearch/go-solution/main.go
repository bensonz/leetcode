package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	// basically just a dfs
	return true
}

func main() {
	a := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(a)
	fmt.Println(exist(a, "ABCCED"))
}
