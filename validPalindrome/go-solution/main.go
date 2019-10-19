package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	var i, j int
	i = 0
	j = len(s) - 1
	for i < j {
		for (i < j) && !((s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')) {
			i += 1
		}

		for (i < j) && !((s[j] >= '0' && s[j] <= '9') || (s[j] >= 'a' && s[j] <= 'z') || (s[j] >= 'A' && s[j] <= 'Z')) {
			j -= 1
		}

		if i < j && s[i] != s[j] {
			return false
		}

		i++
		j--
	}
	return true
}

func main() {
	a := "aAza"
	fmt.Println(isPalindrome(a))

}
