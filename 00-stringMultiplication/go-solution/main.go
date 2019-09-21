package main

import (
	"fmt"
)

/*
This is actually pretty weird.
Something like i + j determines the i+j
*/
func getMaxSize() int {
	MaxInt := 1 << 31
	count := 0
	for MaxInt > 0 {
		MaxInt /= 10
		count++
	}
	return count
}

func multiply(num1 string, num2 string) string {

	return ""
}

func main() {
	fmt.Println(multiply("0", "13241234887031495"))
	fmt.Println(multiply("123", "4567"))
}
