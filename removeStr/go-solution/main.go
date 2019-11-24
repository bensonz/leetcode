package main

import (
	"fmt"
)

// removes char in s2 from s1
// i.e  s1 abcdee
//      s2 aaa
// return bcdee
func remove(s1, s2 string) string {
	var mapping [26]bool
	for _, c := range s2 {
		mapping[c-'a'] = true
	}
	result := ""
	for i, c := range s1 {
		if mapping[c-'a'] {
			// remove this
		} else {
			// add to our string
			result += s1[i : i+1]
		}
	}
	return result
}

func main() {
	fmt.Println(remove("abcdeeff", "af"))

	fmt.Println(remove("abcdeeff", "a"))

	fmt.Println(remove("abcdeeff", ""))

	fmt.Println(remove("abcdeeff", "zcxvzx"))
}
