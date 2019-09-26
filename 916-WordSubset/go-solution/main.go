package main

import (
	"fmt"
)

var primeNumbers = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func hashWord(s string) int {
	hash := 1
	var idx int
	for _, sr := range s {
		idx = int(sr - 'a')
		hash *= primeNumbers[idx]
	}
	// if hash < 1, this string is empty
	return hash
}

func hashArray(B []string) map[int]int {
	var result = make(map[int]int)
	var hw int
	for _, bs := range B {
		hw = hashWord(bs)
		result[hw] = 1
	}
	return result
}

func checkWord(toCheck string, bHash map[int]int) bool {
	ah := hashWord(toCheck)
	for key := range bHash {
		if (ah % key) != 0 {
			// not divisible! a has different number of chars
			return false
		}
	}
	return true
}

func wordSubsets(A []string, B []string) []string {
	// var result []string
	iteration, count := 0, 0 // this keeps track of our stuff
	aLength := len(A)
	bHash := hashArray(B)
	fmt.Println(bHash)
	for (iteration + count) < aLength {
		toCheck := A[iteration]
		if checkWord(toCheck, bHash) {
			// pass
			fmt.Printf("%v passed.\n", toCheck)
			iteration++
		} else {
			// swap with last element
			A[iteration], A[aLength-count-1] = A[aLength-count-1], A[iteration]
			count++
		}
	}
	return A[0:iteration]
}

func main() {
	a := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	var b []string

	b = []string{"e", "o"}
	fmt.Println(wordSubsets(a, b))
	// b = []string{"e", "oo"}
	// fmt.Println(wordSubsets(a, b))
	// b = []string{"lo", "eo"}
	// fmt.Println(wordSubsets(a, b))
}
