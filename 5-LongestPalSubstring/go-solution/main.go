package main

import (
	"fmt"
)

/*
obviously there's brute force, loop through it and find the length of every palindrome.
*/

func extendPalindrome(s string, i, j int) (length, start int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return j - i - 1, i + 1
}

func longestPalindrome(s string) string {
	sLength := len(s)
	if sLength < 2 {
		return s
	}
	max := 0
	start := 0
	for i := 0; i < len(s); i++ {
		lo, so := extendPalindrome(s, i, i)
		// odd
		if max < lo {
			max = lo
			start = so
		}
		le, se := extendPalindrome(s, i, i+1)
		// even
		if max < le {
			max = le
			start = se
		}
	}
	return s[start : start+max]
}

func main() {
	l := longestPalindrome
	fmt.Println(l(""))
	fmt.Println(l("ac"))
	fmt.Println(l("bb"))
	fmt.Println(l("babad"))
	fmt.Println(l("cbbddbb"))
	fmt.Println(l("abcdeffedjklkjd"))
}
