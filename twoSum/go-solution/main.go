package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		s := nums[i] + nums[j]
		if s == target {
			break
		}
		if s > target {
			j--
		}
		if s < target {
			i++
		}
	}
	return []int{i, j}
}

func main() {
	a := []int{2, 7, 12, 4, 5, 0, 1}
	fmt.Println(twoSum(a, 9))
}
