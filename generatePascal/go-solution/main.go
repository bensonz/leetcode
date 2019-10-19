package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	var result = make([][]int, numRows)
	if numRows < 1 {
		return result
	}
	result[0] = []int{1} // base row
	// so we can start loop at second row
	for i := 1; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}

func main() {
	g := generate(5)
	for i := 0; i < len(g); i++ {
		fmt.Println(g[i])
	}
}
