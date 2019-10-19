package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func twoSumWithClosest(arr []int, k int) (int, int) {
	left := 0
	right := len(arr) - 1
	ansLeft, ansRight := 0, 0
	diff := 1 << 31
	for left < right {
		off := arr[left] + arr[right] - k
		if off == 0 {
			// found an exact match
			return left, right
		}
		if abs(off) < diff {
			// store result
			diff = off
			ansLeft = left
			ansRight = right
		} else if off < 0 {
			left++
		} else if off > 0 {
			right--
		}
	}
	return ansLeft, ansRight
}

func main() {
	a := []int{1, 2, 5, 7, 10, 14, 99}
	t := twoSumWithClosest
	fmt.Println(t(a, 99))
	fmt.Println(t(a, 100))
	fmt.Println(t(a, 9))
}
