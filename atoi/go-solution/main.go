package main

import (
	"fmt"
)

func myAtoi(str string) int {
	rray := []int{}
	sign := false
	i := 0
	for {
		if i >= len(str) {
			return 0
		}
		checkRune := str[i]
		if checkRune == ' ' {
			i++
		} else if checkRune == '-' {
			i++
			sign = true
			break
		} else if checkRune == '+' {
			i++
			break
		} else if checkRune <= '9' && checkRune > '0' {
			break
		} else if checkRune == '0' {
			i++
		} else {
			// didnt start correctly
			return 0
		}
	}

	for i < len(str) {
		if !(str[i] >= '0' && str[i] <= '9') {
			break
		}
		rray = append(rray, int(str[i]-'0'))
		i++
	}
	l := len(rray)
	fmt.Println(rray)
	if len(rray) == 0 {
		return 0
	}
	c := 1
	iMax := 1 << 31
	result := rray[l-1]
	for j := l - 2; j >= 0; j-- {
		result += rray[j] * (10 * c)
		c *= 10
		if c >= (100000000) || result > iMax {
			if sign {
				return -iMax
			}
			return iMax - 1
		}
	}

	if sign {
		return -result
	}
	return result
}
func main() {
	fmt.Println(1 << 31)
	fmt.Println(myAtoi("2147483231412648"))
}
