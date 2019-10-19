package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	l := len(nums)
	dp := make([]int, l)

	if l <= 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	}
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[l-1]
}

func main() {
	a := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(a))
}
