package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	l := len(digits)
	v := digits[l-1]
	if v+1 == 10 {
		var i int
		for i = l - 1; i >= 0; i-- {
			if digits[i] != 9 {
				break
			} else {
				// carry
				digits[i] = 0
			}
		}
		if i == -1 {
			digits = append([]int{1}, digits...)
		} else {
			digits[i]++
		}
	} else {
		digits[l-1] = v + 1
	}
	return digits
}
func main() {
	fmt.Println(plusOne([]int{9, 9}))
}
