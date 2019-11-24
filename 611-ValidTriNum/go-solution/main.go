package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	nlen := len(nums)
	count := 0
	for i := nlen - 1; i > 1; i-- {
		j, k := 0, i-1
		for j < k {
			if nums[j]+nums[k] > nums[i] {
				count += k - j
				k--
			} else {
				j++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}
