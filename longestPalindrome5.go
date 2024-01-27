package main

import (
	"fmt"
)

func main() {
	str := "babad"
	fmt.Println(longestPalindrome(str))
}
func longestPalindrome(s string) string {
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		l1 := getMaxLen(s, i, i)
		l2 := getMaxLen(s, i, i+1)
		maxLength := l1
		if l2 > l1 {
			maxLength = l2
		}
		if maxLength > (end - start + 1) {
			start = i - (maxLength-1)/2
			end = i + (maxLength)/2
		}
	}

	return s[start : end+1]
}

func getMaxLen(s string, left int, right int) int {
	if len(s) == 0 || left > right {
		return 0
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return (right - 1) - (left + 1) + 1
}
