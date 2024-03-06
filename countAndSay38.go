package main

import (
	"fmt"
	"strconv"
)

func main() {

	println(countAndSay(9))

}

func countAndSay(n int) string {
	result := make([]string, n)
	result[0] = "1"

	for i := 1; i < n; i++ {
		result[i] = getSay(result[i-1])
	}
	fmt.Println(result)
	return result[n-1]
}

func getSay(str string) string {
	s := ""
	l := len(str)
	println(s, l, str)
	for i := 0; i < l; i++ {
		count := 1
		for i < (l-1) && str[i] == str[i+1] {
			count++
			i++
		}
		s = s + strconv.Itoa(count) + string(str[i])
	}
	return s
}
