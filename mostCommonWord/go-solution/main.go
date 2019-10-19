package main

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)
	bannedMap := make(map[string]bool)
	for _, w := range banned {
		bannedMap[w] = true
	}
	freq := make(map[string]int)

	i, l := 0, len(paragraph)
	for i < l {
		for i < l && (paragraph[i] < 'a' || paragraph[i] > 'z') {
			i++
		}
		i1 := i
		for i < l && paragraph[i] >= 'a' && paragraph[i] <= 'z' {
			i++
		}
		w := string(paragraph[i1:i])
		if bannedMap[w] {
			continue
		}
		freq[w]++
	}
	str := ""
	count := 0
	for w, f := range freq {
		if f > count {
			str, count = w, f
		}
	}
	return str
}

func main() {
	m := mostCommonWord
	p := "Bob hit a ball, the hit BALL flew far after it was hit."
	fmt.Println(m(p, []string{"hit"}))
}
