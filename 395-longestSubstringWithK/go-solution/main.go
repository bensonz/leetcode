package main

import (
	"fmt"
)

func validate(m map[rune]int, k int) bool {
	for _, val := range m {
		if val < k {
			return false
		}
	}
	return true
}

func longestSubstring(s string, k int) int {
	alphabet := make(map[rune]int)
	for _, char := range s {
		alphabet[char]++
	}
	fmt.Println(alphabet)
	max, cur := 0, 0
	for i := range s {

	}
	return max
}

func main() {
	f := longestSubstring
	fmt.Println(f("ababbc", 2))
	fmt.Println(f("abbcabbefffc", 2))
}
