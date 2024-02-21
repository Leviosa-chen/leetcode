package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i <= l/2-1; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}
