package main

import (
	"fmt"
	"strconv"
)

func flatten(s string) string {
	// no [ or ] inbetween start and end
	start, end := -1, -1
	for i := range s {
		if s[i] == '[' {
			start = i
		} else if s[i] == ']' {
			end = i
		}
	}
	result := ""
	numberTail := 1
	d := s[start-numberTail]
	for numberTail < start && d >= '0' && d <= '9' {
		numberTail++
		d = s[start-numberTail]
	}

	mul, _ := strconv.Atoi(s[start-numberTail : start])
	// fmt.Printf("Start: %v | End: %v | NumberTail : %v | Mul : %v\n", start, end, numberTail, mul)
	for i := 0; i < mul; i++ {
		result += s[start+1 : end]
	}
	return result
}

func decodeString(s string) string {
	stack := []int{}
	l := len(s)
	if l < 1 {
		// nothing to do
		return ""
	}
	result := ""
	for i := range s {
		if s[i] == '[' {
			stack = append(stack, i)
		} else if s[i] == ']' {
			if len(stack) > 0 {
				// pop last element
				lastIdx := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				fmt.Println(lastIdx)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(flatten("90[abc]"))
	fmt.Println(flatten("9[abc]"))
	// fmt.Println(decodeString("3[a]3[bbb280[z]]efg"))
}
