package main

import (
	"fmt"
)

// LPS find the longest prefix and suffix of a string
func LPS(s string) []int {
	l := len(s)
	result := make([]int, l)
	i, j := 1, 0
	for i < l {
		fmt.Printf("Checking i: %d, j: %d\n", i, j)
		if s[i] == s[j] {
			// match, move both
			// record down
			j++
			result[i] = j
			i++
			fmt.Printf("Match: %v\n", result)
		} else if j > 0 {
			j = result[j-1]
			fmt.Printf("No Match Setback j: %v\n", result)
		} else {
			i++
		}
	}
	return result
}

func main() {
	fmt.Println(LPS("aaabbbaaa"))
}
