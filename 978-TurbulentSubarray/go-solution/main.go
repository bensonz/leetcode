package main

import (
	"fmt"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxTurbulenceSize(A []int) int {
	l := len(A)
	if l <= 1 {
		return l
	}
	greater, less := 1, 1
	res := 1
	for i := 1; i < l; i++ {
		switch {
		case A[i] > A[i-1]:
			greater, less = less+1, 1
			res = max(res, greater)
		case A[i] < A[i-1]:
			greater, less = 1, greater+1
			res = max(res, less)
		default:
			greater, less = 1, 1
		}
	}

	return res
}
func main() {
	m := maxTurbulenceSize
	// fmt.Println(m([]int{9, 4, 2, 10, 7, 8, 8, 1, 9})) // 5
	b := []int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24}
	fmt.Println(m(b)) // 8
	// c := []int{100, 100, 100}
	// fmt.Println(m(c)) // 1
}
