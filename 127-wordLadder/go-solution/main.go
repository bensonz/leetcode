package main

import (
	"fmt"
)

func dOne(a, b string) bool {
	// a and b are of same length
	// fmt.Printf("Checking: %v | %v :", a, b)
	d := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
			if d > 1 {
				// fmt.Printf("False\n")
				return false
			}
		}
	}
	// fmt.Printf("True\n")
	return d == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var nextQueue, tmpQueue, leftOver []string
	var count int
	var used = make(map[string]bool)

	nextQueue = append(nextQueue, beginWord)
	for len(nextQueue) > 0 {
		for _, e := range nextQueue {
			for _, word := range wordList {
				if used[word] || word == e {
					// already did something with it
					continue
				}
				if dOne(e, word) {
					// add to tmpQueue, gonna check this elem next time
					if word == endWord {
						return count + 2
					}
					used[word] = true
					tmpQueue = append(tmpQueue, word)
				} else {
					// add to leftOver
					used[word] = true
					leftOver = append(leftOver, word)
				}
			}
		}
		// increment count, next iteration
		count++
		fmt.Println(count)
		fmt.Println("NextQ:", tmpQueue)
		fmt.Println("LeftO:", leftOver)
		// set variables
		nextQueue = tmpQueue
		wordList = leftOver
		used = make(map[string]bool)
		// reset
		leftOver = []string{}
		tmpQueue = []string{}
	}
	return 0
}

func main() {
	b, e := "hit", "cog"
	wl := []string{"hot", "dog", "dot", "lot", "log"}
	fmt.Println(ladderLength(b, e, wl))
}
