package main

import (
	"fmt"
)

// Min ...
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/*
 solution idea
  This question is pretty interesting

  First Imma solve this with the most staright forward way

  Here's result :
  Runtime: 280 ms, faster than 29.98% of Go online submissions for Container With Most Water.
  Memory Usage: 5.6 MB, less than 25.74% of Go online submissions for Container With Most Water.
*/
// func maxArea(height []int) int {
// 	var max int
// 	for i := 0; i < len(height); i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			h1 := height[i]
// 			h2 := height[j]
// 			area := Min(h1, h2) * (j - i)
// 			if area > max {
// 				max = area
// 			}
// 		}
// 	}
// 	return max
// }

/*
  but let's try to do better. One possible idea is to use two pointers. As we know that when the distance between them shrinks, whichever height is larger will produce the larger area.
  So the problem becomes, who are you gonna move, the left pointer or the right pointer?
  So obviously you'd want to keep the higher one, so you just move the lower one.

  How do we prove that we have the max area?

  At first, we get the max area, no problem with that.

  Then we move a single pointer, reducing its distance by 1, and the result area is the max area possible, as if you have moved the other pointer, with the same distance, you'd have smaller area.
  So we picked the largest with current condition.

  Heres' the code:
*/
func calArea(height []int, i, j int) int {
	h1 := height[i]
	h2 := height[j]
	return Min(h1, h2) * (j - i)
}

func maxArea(height []int) int {
	var max int
	start := 0
	end := len(height) - 1
	for start < end {
		area := calArea(height, start, end)
		if area > max {
			max = area
		}
		if height[start] < height[end] {
			// move start
			start++
		} else {
			// move end
			end--
		}
	}
	return max
}
func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(a) == 49)

	b := []int{1, 2, 3, 4}
	fmt.Println(maxArea(b) == 4)

	c := []int{100, 100, 1, 1, 1}
	fmt.Println(maxArea(c) == 100)
}
