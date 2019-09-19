package main

import (
	"fmt"
	"sort"
)

func find(nums []int, t int, c chan int) {
	for _, v := range nums {
		if v == t {
			c <- v
			return
		}
	}
	c <- -1
	return
}

func threeSum(nums []int) [][]int {

}

func main() {
	nums := []int{-1, 0, 1, 2, -1, 4}
	fmt.Println(threeSum(nums))
}
