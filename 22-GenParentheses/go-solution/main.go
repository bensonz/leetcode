package main

import (
	"fmt"
)

func gPar(s string, max, left, right int, ans *[]string) {
	if len(s) == max*2 {
		// got one
		*ans = append(*ans, s)
		return
	}
	if right > left {
		// unmatch
		return
	}
	// otherwise branch
	if left < max {
		gPar(s+"(", max, left+1, right, ans)
	}
	if right < max {
		gPar(s+")", max, left, right+1, ans)
	}
}

func generateParenthesis(n int) []string {
	var result []string
	if n <= 0 {
		return result
	}
	gPar("", n, 0, 0, &result)
	return result
}

func main() {
	a := generateParenthesis(3)
	for _, row := range a {
		fmt.Println(row)
	}
}
