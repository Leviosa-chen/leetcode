package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(addBinary("1010", "1011"))

}

func addBinary(a string, b string) string {
	la := len(a) - 1
	lb := len(b) - 1
	add := 0

	result := []string{}
	for la >= 0 || lb >= 0 {
		m := 0
		if la >= 0 {
			m, _ = strconv.Atoi(string(a[la]))
		}
		n := 0
		if lb >= 0 {
			n, _ = strconv.Atoi(string(b[lb]))
		}
		plus := m + n + add
		if plus >= 2 {
			add = 1
		} else {
			add = 0
		}
		result = append(result, strconv.Itoa(plus%2))
		println(plus, m, n, add)
		la--
		lb--
	}
	if add == 1 {
		result = append(result, "1")
	}
	revertArr(result)
	return strings.Join(result, "")
}

func revertArr(digits []string) []string {
	l := len(digits)
	for i := 0; i < l/2; i++ {
		digits[i], digits[l-1-i] = digits[l-1-i], digits[i]
	}
	return digits
}
