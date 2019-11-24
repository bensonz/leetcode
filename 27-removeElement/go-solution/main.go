package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	if j < 1 {
		return 0
	}
	for i <= j {
		if nums[i] == val {
			// need remove this, swap
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	fmt.Println(nums)
	nums = nums[0:i]
	fmt.Println(nums)
	return i
}

func main() {
	a := []int{3, 2, 2, 3}
	fmt.Println(removeElement(a, 3))
	b := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(b, 2))
	c := []int{3, 3}
	fmt.Println(removeElement(c, 5))
}
