package main

import (
	"fmt"
	"reflect"
)

/*
  So this asks for find any permutation of a string in another string.
  idea 1: (The obvious idea)
  list all permutations of string 1, then check for each permutation
  Runtime:
  permutation: O(2^n) n is the length of string 1
  Checking for string: O(m) m is the length of string 2
  Total = O(2^n * n)
  which is extremely slow.
  idea 2:
  Well, since all we need to do is to find a permutation of s1 in s2, we can focus on the number of each characters. For example, if s1 is "abb", we can say it contains 1 a, and 2 b, putting it into a data structure would be similar to a map {a:1, b:2}
  Then our goal is to check whether s2 contains this structure. If we simply check for the whole string, s2 could contain such structure but these characters can be scattered throughout the string. Thus enter our concept of a sliding window.
*/
func transput(s string) map[rune]int {
	var result = make(map[rune]int)
	for _, e := range s {
		count := result[e]
		if count != 0 {
			result[e] = count + 1
		} else {
			result[e] = 1
		}
	}
	return result
}

func checkInclusion(s1 string, s2 string) bool {
	var windowSize = len(s1)
	var translate = transput(s1)
	var myRange = len(s2) - windowSize
	for i := 0; i <= myRange; i++ {
		fmt.Printf("Chekcing %s \n", s2[i:i+windowSize])
		s2Translate := transput(s2[i : i+windowSize])
		if reflect.DeepEqual(translate, s2Translate) {
			return true
		}
	}
	return false
}

func c(s1 string, s2 string) bool {
	return checkInclusion(s1, s2)
}

func main() {
	fmt.Println(c("adc", "dcda"))
	// fmt.Println(c("ab", "eidbaooo"))
	// fmt.Println(c("ab", "eidbooo"))
	// fmt.Println(c("ab", ""))
	// fmt.Println(c("ab", "baababba"))
	// fmt.Println(c("", ""))
	// fmt.Println(c("abcd", "zbcdazxeq"))
}
