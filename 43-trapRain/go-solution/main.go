package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {

}

func main() {
	// a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// fmt.Println(trap(a))
	// b := []int{5, 4, 1, 2}
	// fmt.Println(trap(b))
	// c := []int{5, 2, 1, 2, 1, 5}
	// fmt.Println(trap(c))
	d := []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}
	fmt.Println(trap(d))
}
