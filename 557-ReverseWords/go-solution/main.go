package main

import (
	"fmt"
)

func reverseWord(s string) string {
	var result string
	for i := len(s) - 1; i >= 0; i-- {
		result = fmt.Sprintf("%s%c", result, s[i])
	}
	// fmt.Printf("Original:|%s| Rev:|%s|\n", s, result)
	return result
}

func reverseWords(s string) string {
	var result string
	lastI := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == ' ' {
			r := reverseWord(s[lastI:i])
			result = fmt.Sprintf("%s%s ", result, r)
			lastI = i + 1
		}
	}
	// fmt.Println(lastI, len(s))
	// if lastI != len(s)-1 {
	result = fmt.Sprintf("%s%s", result, reverseWord(s[lastI:]))
	// }
	return result
}

func main() {
	a := "let's go to the mall"
	b := "i love u"
	fmt.Println(reverseWords(a))
	fmt.Println(reverseWords(b))
}
