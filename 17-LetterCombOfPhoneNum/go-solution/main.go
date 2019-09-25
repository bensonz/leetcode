package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	phoneMap := map[string][]string{
		"0": []string{},
		"1": []string{},
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	var result []string
	for i := 0; i < len(digits); i++ {
		d := digits[i : i+1]
		m := phoneMap[d]
		// fmt.Println(d, m)
		if len(m) == 0 {
			continue
		}
		if len(result) == 0 {
			// empty, put m in
			result = m
			continue
		}
		var newResult []string
		for _, origin := range result {
			// fmt.Println(origin)
			for _, c := range m {
				// fmt.Println(c)
				e := fmt.Sprintf("%v%v", origin, c)
				// fmt.Println(e)
				newResult = append(newResult, e)
			}
		}
		result = newResult
	}
	return result
}

func main() {
	l := letterCombinations
	fmt.Println(l("9013"))
	// fmt.Println(l("23"))
}
