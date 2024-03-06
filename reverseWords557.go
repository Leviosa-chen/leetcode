package main

import "fmt"

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	bytes := []byte(s)
	l := len(bytes)
	for i := 0; i < l; i++ {
		start := i
		for i < l && string(bytes[i]) != " " {
			println(string(bytes[i]), i)
			i++
		}
		end := i - 1
		for start < end {
			bytes[start], bytes[end] = bytes[end], bytes[start]
			start++
			end--
		}
	}
	return string(bytes)
}
