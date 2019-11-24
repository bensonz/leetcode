package main

import (
	"fmt"
)

func longestSubarrayWithInc(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n, max := len(nums), 1
	d := make([]int, n)
	d[0] = 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			cur := 1
			if nums[i] > nums[j] {
				cur = d[j] + 1
			}
			d[i] = Max(d[i], cur)
		}
		max = Max(max, d[i])
	}
	return max
}

// Max ...
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	l := longestSubarrayWithInc
	a := []int{1, 5, 8, 0, 2, 3, 100, 4, 5, 3, 16}
	fmt.Println(l(a))
}
