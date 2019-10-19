package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k != 0 {
		copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	}
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	fmt.Println(a)
	a = []int{1, 2}
	rotate(a, 0)
	fmt.Println(a)
}
