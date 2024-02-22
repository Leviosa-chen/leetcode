package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiply("2", "3"))
}

// TODO 未完成
func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)

	mult := 0
	for i := l1 - 1; i >= 0; i-- {
		tmp_mult := 0
		add := 0
		for j := l2 - 1; j >= 0; j-- {
			a := int(num1[i])*int(num2[j]) + add
			add = a / 10
			tmp_mult += (a % 10)
		}
		mult += tmp_mult
	}
	return string(mult)
}
