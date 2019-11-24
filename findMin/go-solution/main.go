package main

import (
	"fmt"
)

// returns the index of the minium number in an array
func findMinInRotatedArray(array []int) int {
	left, right := 0, len(array)-1
	if right == -1 {
		return -1
	}
	if right == 0 {
		return 0
	}

	mid := left + (right-left)/2
	for mid >= 0 && mid < len(array) {
		leftSort := array[mid] > array[left]
		rightSort := array[mid] < array[right]
		if leftSort && rightSort {
			// ok, this part is compleletly sorted
			return left
		}
		if leftSort {
			// right is not sorted, our element is in right, move left
			left = mid + 1
		}
		if rightSort {
			right = mid
		}
		mid = left + (right-left)/2
	}
	return -1
}

func main() {
	a := []int{50, 52, 63, 90, 3, 8, 15, 44}
	fmt.Println(findMinInRotatedArray(a))
	b := []int{1, 2, 3, 4, 5}
	fmt.Println(findMinInRotatedArray(b))
}
