package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(str, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	mod := numRows*2 - 2
	arr := make([]string, numRows)

	for i, char := range s {
		x := i % mod
		index := min(x, mod-x)
		arr[index] += string(char)
	}

	return strings.Join(arr, "")
}
