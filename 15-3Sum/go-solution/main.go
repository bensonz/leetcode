package main

import (
	"fmt"
	"sort"
)

// target is 0
func threeSum(nums []int) [][]int {
	nlen = len(nums)
	sort.Ints(nums) // sort it
	var result [][]int
	if nlen < 3 || nums[0] > 0 || nums[nlen-1] < 0 {
		// no solution
		return result
	}
	for i := 0; i <= nlen-3; i++ {
		j := i + 1
		k := nlen - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for ; j < k && nums[j] == nums[j-1]; j++ {

				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, 4}
	fmt.Println(threeSum(nums))
}
