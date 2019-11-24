package main

import (
	"fmt"
)

func checkLegal(history []bool, idx int) (int, bool) {
	// we know history[idx] is true, he drove at this day
	distance := 0
	for i := 3; i > 0; i-- {
		// count the next four days of driven times
		if idx+i > len(history)-1 {
			// end
			return len(history) - 1, true
		}
		if history[idx+i] == true {
			distance = i
			break
		}
	}
	idx += 3 // move forward 3 days
	for j := 1; j < distance+2; j++ {
		if history[idx+j] != false {
			// not ok
			return len(history) - 1, false
		}
	}
	return idx + distance, true
}

func isLegal(history []bool) bool {
	for i := 0; i < len(history); i++ {
		if history[i] == true {
			// did drive, fit the next 8 days
			nextI, legal := checkLegal(history, i)
			if !legal {
				return false
			}
			i = nextI
		}
	}
	return true
}

func main() {
	a := []bool{true, false, false, false, false, true}
	fmt.Println(isLegal(a))
	b := []bool{true, true, false, false, true, false}
	fmt.Println(isLegal(b))
	fmt.Println(isLegal(append(a, b...)))
}
