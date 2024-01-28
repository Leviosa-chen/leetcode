package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	result := 0
	for x != 0 {
		if temp := int32(result); (temp*10)/10 != temp {
			return 0
		}
		result = result*10 + x%10
		x = x / 10
	}
	return result
}
