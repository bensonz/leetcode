package main

import (
	"fmt"
)

// question link:
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

/*
soluction idea:

so we are looking for the longest substring here, we need to find all
non-repeating words first, then see which has the longest length and return that
length, which we could probably do with map()

*/

// func readUntilRepeat(s string) string {
// 	cur := ""
// 	myMap := make(map[rune]int32)
// 	for _, c := range s {
// 		_, ok := myMap[c]
// 		if ok {
// 			// seen
// 			break
// 		} else {
// 			myMap[c] = 1
// 			cur += string(c)
// 		}
// 	}
// 	fmt.Printf("Until repeat: %s\n", cur)
// 	return cur
// }
//
// func lengthOfLongestSubstring(s string) int {
// 	longestSubstring := ""
// 	cur := ""
// 	for idx := range s {
// 		cur = readUntilRepeat(s[idx:])
// 		if len(cur) > len(longestSubstring) {
// 			longestSubstring = cur
// 		}
// 	}
// 	fmt.Printf("LongestSubstring: %s\n", longestSubstring)
// 	return len(longestSubstring)
// }
func lengthOfLongestSubstring(s string) int {
	max := 0
	begin := 0
	hash := make(map[rune]int)
	for i, c := range s {
		index, ok := hash[c]
		length := i - begin + 1 // from current index to the begin
		if ok && index >= begin {
			// if repeated, and the seen char index is greater than begin
			// we update the length and begin
			length = i - index
			begin = index + 1
		}
		hash[c] = i
		fmt.Printf("length : %d, max: %d \n", length, max)
		if length > max {
			max = length
		}
	}

	return max
}

func l(s string) int {
	return lengthOfLongestSubstring(s)
}

func main() {
	fmt.Println(l("abcabcabc")) // abc
	fmt.Println(l("bbbbbb"))    // b
	fmt.Println(l("13414be"))   // 14be
	fmt.Println(l("pwwkew"))    // wke
	fmt.Println(l("dvdf"))      // vdf

}
