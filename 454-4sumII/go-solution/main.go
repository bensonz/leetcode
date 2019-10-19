package main

import (
	"fmt"
)

func fourSumCount(A []int, B []int, C []int, D []int) int {
	AB := make(map[int]int)
	var sum int
	for _, a := range A {
		for _, b := range B {
			sum = a + b
			AB[sum]++
		}
	}
	CD := make(map[int]int)
	for _, c := range C {
		for _, d := range D {
			sum = c + d
			CD[sum]++
		}
	}
	final := 0
	for ab, occur := range AB {
		if occur2, ok := CD[-ab]; ok {
			final += occur * occur2
		}
	}
	return final
}

func main() {
	f := fourSumCount
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(f(A, B, C, D))
}
