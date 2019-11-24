package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	l := len(nums)
	var front, back, ans []int = make([]int, l), make([]int, l), make([]int, l)
	fcoeff, bcoeff := nums[0], nums[l-1]
	front[0], back[l-1] = 1, 1 // for multiplication later
	// fmt.Println(fcoeff, bcoeff)
	for i := 1; i < l; i++ {
		front[i] = fcoeff
		fcoeff *= nums[i]

		back[l-i-1] = bcoeff
		bcoeff *= nums[l-i-1]
	}
	// fmt.Println(front, back)
	for j := 0; j < len(front); j++ {
		ans[j] = front[j] * back[j]
	}
	return ans
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(a))
}
