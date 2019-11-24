package main

import (
	"fmt"
)

func permHelp(nums []int, used []bool, result *[][]int, temp []int) {
	if len(temp) == len(nums) {
		r := append([]int{}, temp...) // make a copy of temp
		*result = append(*result, r)
		return
	}
	for i, num := range nums {
		if used[i] {
			continue
		}
		used[i] = true
		// add this num to our temp
		temp = append(temp, num)
		// then permutate the rest
		permHelp(nums, used, result, temp)
		temp = temp[:len(temp)-1]
		used[i] = false
	}
}

func permutate(nums []int) [][]int {
	var result [][]int
	if len(nums) < 1 {
		return result
	}
	used := make([]bool, len(nums))
	temp := []int{}
	permHelp(nums, used, &result, temp)
	return result
}

func main() {
	c := permutate([]int{1, 2, 3})
	fmt.Println(c)
}
