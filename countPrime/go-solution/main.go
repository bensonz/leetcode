package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	root := int(math.Pow(float64(n), 0.5))
	for i := 2; i <= root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	count := 1
	for i := 3; i < n; i += 2 {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(12))
}
