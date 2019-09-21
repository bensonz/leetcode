package main

import (
	"fmt"
)

// returns the index
func bSearch(target int, slice []int) int {
	l := len(slice)
	if l < 1 {
		return -1 // not here
	}
	if l == 1 {
		if slice[0] == target {
			return 1
		}
		return -1
	}
	// otherwise search
	start := l / 2
	elem := slice[start]
	if elem == target {
		// found
		return start
	} else if elem < target {
		return start + bSearch(target, slice[start:])
	} else {
		return bSearch(target, slice[0:start])
	}
}

// this one doesn't create the slices, uses indexes instead
func binarySearch(target int, slice []int) int {
	left := 0
	right := len(slice) - 1
	for left <= right {
		middle := left + (right-left)/2
		if slice[middle] == target {
			return middle
		} else if slice[middle] < target {
			// search right
			left = middle + 1
		} else {
			//search left
			right = middle - 1
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 8, 10, 13, 16, 19, 20, 25}
	fmt.Println(bSearch(13, a))
	fmt.Println(binarySearch(13, a))
	fmt.Println(bSearch(0, a))
	fmt.Println(binarySearch(0, a))
}
