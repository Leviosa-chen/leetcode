package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	letters := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	result := []string{}
	for i := 0; i < len(digits); i++ {
		char := string(digits[i])

		if i == 0 {
			for _, letter := range letters[char] {
				result = append(result, letter)
			}
		} else {
			tmpResult := []string{}
			for _, letter := range letters[char] {
				for _, res := range result {
					tmpResult = append(tmpResult, res+letter)
				}
			}
			result = tmpResult
		}
	}

	return result
}
