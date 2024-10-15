package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
}

func plusOne(digits []int) []int {

	digits = revert(digits)
	add := 0
	for i := 0; i < len(digits); i++ {
		add = digits[i] + 1
		if add == 10 {
			digits[i] = add - 10
		} else {
			digits[i] = add
			break
		}
	}
	if add == 10 {
		digits = append(digits, 1)
	}
	return revert(digits)
}

func revert(digits []int) []int {
	l := len(digits)
	for i := 0; i < l/2; i++ {
		digits[i], digits[l-1-i] = digits[l-1-i], digits[i]
	}
	return digits
}
