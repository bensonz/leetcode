package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		// nothing here
		return 0
	}
	left, right, sum, count := 0, 0, 0, 0
	for right = range nums {
		sum += nums[right]

		fmt.Printf("Moving Right:%v, Sum : %v \n", right, sum)
		for sum == k {
			count++
			fmt.Println("Moving left in equal:", left)
			sum -= nums[left]
			left++
		}
		// move left to see if we may hold
		for left <= right && sum != k {
			fmt.Println("Moving left in neq :", left)
			sum -= nums[left]
			left++
		}
	}
	return count
}

func main() {
	a := []int{1, 1, 1}
	fmt.Println(subarraySum(a, 2))
}
