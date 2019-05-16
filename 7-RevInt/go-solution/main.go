package main

import (
	"fmt"
)

func reverseHelp(x int32) int32 {
	var result, m, n int32
	max := int32((1 << 31) / 10)
	min := -max
	for {
		m = x / 10
		n = x % 10
		if m == 0 {
			// less than 10
			if n != 10 {
				return result*10 + n
			}
		}
		result = result*10 + n
		if result < min || result > max {
			// going to overflow
			return 0
		}
		x = m
	}
}

func reverse(x int) int {
	k := int32(x)
	y := reverseHelp(k)
	return int(y)
}

func main() {
	fmt.Println(reverse(123) == 321)
	fmt.Println(reverse(0) == 0)
	fmt.Println(reverse(10) == 1)
	fmt.Println(reverse(1) == 1)
	fmt.Println(reverse(-1) == -1)
	fmt.Println(reverse(-123) == -321)
	fmt.Println(reverse(1534236469) == 0)
	fmt.Println(reverse(-2147483648) == 0)
	fmt.Println(reverse(1563847412) == 0)
	// max = 4294967295
	// myOutput = 2147483651
}
