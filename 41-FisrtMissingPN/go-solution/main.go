package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	cp := make(map[int]int)
	for _, k := range nums {
		if k <= 0 {
			continue
		}
		// otherwise put a 1 at the corresponding position
		cp[k] = 1
	}
	i := 1
	for {
		if cp[i] == 1 {
			// exist
			i++
		} else {
			break
		}
	}
	return i
}

func main() {
	f := firstMissingPositive
	var n []int
	n = []int{0, 1, 2, 3, 6, 7, 10}
	fmt.Println(f(n) == 4)
	n = []int{2, 3, 6, 7, 10}
	fmt.Println(f(n) == 1)
	n = []int{0, 2, 2, 1, 1}
	fmt.Println(f(n) == 3)
}
