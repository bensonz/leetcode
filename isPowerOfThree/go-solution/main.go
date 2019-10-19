package main

import (
	"fmt"
)

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}
	for n%3 == 0 {
		if n == 3 {
			return true
		}
		n /= 3
	}
	return false
}

func main() {
	i := isPowerOfThree
	fmt.Println(i(3 * 3 * 3 * 3 * 4))
}
