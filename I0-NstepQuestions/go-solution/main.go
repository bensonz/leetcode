package main

import (
	"fmt"
)

// this is a solution for having 1 and 2 steps respectively
var baseCase = []int{0, 1, 2}

func numWays(nSteps int) int {
	if nSteps <= 2 {
		return baseCase[nSteps]
	}
	return numWays(nSteps-1) + numWays(nSteps-2)
}

// this is a solution to having x steps respectively
// assuming x is sorted
// func numWays2Help(n int, x []int, xDict map[int]int) int {
// 	if n <= 0 {
// 		// 0 step always 0 ways
// 		return 0
// 	}
// 	if n < x[0] {
// 		// even the smallest cannot finish, quit this
// 		return -1
// 	}
// 	if xDict[x] == 1 {
// 		// has this step, can do
// 		return 1
// 	}
// 	var result int
// 	// otherwise try each step
// 	for single := range x {
// 		rest := n - single
// 		if rest > 0 {
//       //take this step
//       result = result + 1 + numWays2Help(n, x, xDict)
// 		}
// 		if sp != -1 {
// 			result += sp
// 		}
// 	}
// 	return result
// }

func numWays2(n int, x []int) int {
	var result int
	for _, e := range x {
		rest := n - e // this is the rest of steps
		// fmt.Printf("N: %d, Take: %d, Rest:%d\n", n, e, rest)
		if rest > 0 {
			// ok we can take this step
			result += numWays2(rest, x)
		} else if rest == 0 {
			// this way is feasible
			result++
		}
		// otherwise this whole thing fails
	}
	return result
}

func main() {
	// fmt.Println(numWays(3))
	// fmt.Println(numWays(4))
	// fmt.Println(numWays(0))
	// fmt.Println(numWays(16))

	fmt.Println(numWays2(4, []int{1, 3, 5}))
	fmt.Println(numWays2(4, []int{1, 2}))
}
