package main

import (
	"fmt"
)

func traverse(m [][]int, round, row, col int) []int {
	topRow := round
	bottomRow := row - 1 - round
	leftCol := round
	rightCol := col - 1 - round
	var result []int
	if topRow > bottomRow || rightCol < leftCol {
		return result
	}
	// first append top
	for a := leftCol; a <= rightCol; a++ {
		result = append(result, m[topRow][a])
	}
	if topRow == bottomRow {
		return result
	}
	fmt.Println("AfterTop:", result)

	// then right
	for i := topRow + 1; i < bottomRow; i++ {
		result = append(result, m[i][rightCol])
	}
	fmt.Println("AfterRight:", result)
	// then bottom reversed
	for j := rightCol; j >= leftCol; j-- {
		result = append(result, m[bottomRow][j])
	}
	fmt.Println("AfterBottom:", result)
	// then left from bottom to up
	for k := bottomRow - 1; k > topRow; k-- {
		result = append(result, m[k][leftCol])
	}
	fmt.Println("AfterLeft:", result)

	return result
}

func spiralOrder(matrix [][]int) []int {
	var result []int
	if matril == nil || len(matrix) == 0 {
		return result
	}
	row, col := len(matrix), len(matrix[0])
	round := 0
	for {
		fmt.Println("At Round", round)
		tmp := traverse(matrix, round, row, col)
		fmt.Println("Appending:", tmp)
		if len(tmp) == 0 {
			break
		}
		round++
		result = append(result, tmp...)
	}
	return result
}

func main() {
	// a := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// fmt.Println(spiralOrder(a))

	b := [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
	}
	fmt.Println(spiralOrder(b))
}
