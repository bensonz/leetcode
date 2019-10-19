package main

import (
	"fmt"
)

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c && b <= a {
		return b
	}
	if c <= b && c <= a {
		return c
	}
	return b
}

func maximalSquare(matrix [][]byte) int {
	max := 0
	height := len(matrix)
	if height == 0 {
		return 0
	}
	width := len(matrix[0])
	// init
	track := make([][]int, height+1)
	for i := 0; i < height+1; i++ {
		track[i] = make([]int, width+1)
	}
	for i := 1; i < height+1; i++ {
		for j := 1; j < width+1; j++ {
			if matrix[i-1][j-1] == '0' {
				track[i][j] = 0
			} else {
				up := track[i-1][j]
				di := track[i-1][j-1]
				le := track[i][j-1]
				minSq := min(up, di, le) + 1
				track[i][j] = minSq
				if minSq > max {
					max = minSq
				}
			}
		}
	}
	return max * max
}

func main() {
	m := maximalSquare
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(m(matrix))
}
