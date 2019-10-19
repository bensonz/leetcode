package main

import (
	"fmt"
)

func charMatch(a, b byte) bool {
	return a == b || b == '.'
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for row := 0; row < m+1; row++ {
		dp[row] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j < n+1; j++ {
		dp[0][j] = j > 1 && p[j-1] == '*' && dp[0][j-2]
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] != '*' {
				dp[i][j] = dp[i-1][j-1] &&
					charMatch(s[i-1], p[j-1])
			} else {
				dp[i][j] = dp[i][j-2] || dp[i-1][j] &&
					charMatch(s[i-1], p[j-2])
			}
		}
	}
	for row := 0; row < m+1; row++ {
		for col := 0; col < n+1; col++ {
			if dp[row][col] {
				fmt.Printf("1 ") // true print 1
			} else {
				fmt.Print("0 ") // false print 0
			}
		}
		fmt.Println("")
	}
	return dp[m][n]
}

func main() {
	a := "aaabbc"
	p := "c*a*b.."
	fmt.Println(isMatch(a, p))
}
