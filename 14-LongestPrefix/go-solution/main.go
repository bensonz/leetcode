package main

import (
	"fmt"
)

func testAtI(strs []string, i int, prefix string) string {
	for _, s := range strs {
		if s[i:i+1] != prefix {
			return ""
		}
	}
	return prefix
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//just get the shortest so no need to check array bounds later
	shortest := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(shortest) {
			shortest = strs[i]
		}
	}
	var prefix string
	for i := range shortest {
		p := testAtI(strs, i, shortest[i:i+1])
		if p != "" {
			prefix += p
		} else {
			break
		}
	}
	return prefix
}

func main() {
	s := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(s))

	a := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(a))
}
