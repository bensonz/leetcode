package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	var newArr []string
	var result string
	for _, a := range arr {
		if a != "" {
			newArr = append(newArr, a)
		}
	}
	if len(newArr) == 0 {
		return ""
	}
	for i := len(newArr) - 1; i > 0; i-- {
		if newArr[i] == "" {
			continue
		}
		result += newArr[i]
		result += " "
	}
	result += newArr[0]
	return result
}

func main() {
	fmt.Println(reverseWords("  hello world!  "))
}
