package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var con []int
	l1, l2 := len(nums1), len(nums2)
	i, j := 0, 0
	for {
		if i == l1 {
			con = append(con, nums2[j:]...)
			break
		}
		if j == l2 {
			con = append(con, nums1[i:]...)
			break
		}
		v1, v2 := nums1[i], nums2[j]
		if v1 > v2 {
			con = append(con, v2)
			j++
		} else {
			con = append(con, v1)
			i++
		}
	}
	var result float64
	var c = len(con)
	var a = c / 2
	// fmt.Println(con)
	// fmt.Println(c, a)
	if c == 1 {
		result = float64(con[0])
	} else if c%2 == 0 {
		// divisble by two
		result = float64(con[a]+con[a-1]) / 2
		// fmt.Printf("im here: %d, %d = %f\n", con[a], con[a-1], result)
	} else {
		result = float64(con[a])
	}
	return result
}

func f(a, b []int) float64 {
	return findMedianSortedArrays(a, b)
}

func main() {
	z := []int{}
	y := []int{-2, -1}
	a := []int{1, 3}
	b := []int{2}
	c := []int{1, 2}
	d := []int{3, 4}
	fmt.Println(f(z, b)) // 2.0
	fmt.Println(f(a, b)) // 2.0
	fmt.Println(f(c, d)) // (2 + 3) / 2 = 2.5
	fmt.Println(f(a, c)) // (1 + 2) / 2 = 1.5
	fmt.Println(f(a, z)) // (1 + 3) / 2 = 2
	fmt.Println(f(b, y)) //  -1
}
