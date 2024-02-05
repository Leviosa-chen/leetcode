package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(10, 4))
}

func divide(dividend int, divisor int) int {
	resultTmp := dividend / divisor
	if resultTmp < math.MinInt32 {
		return math.MinInt32
	}
	if resultTmp > math.MaxInt32 {
		return math.MaxInt32
	}
	return resultTmp
}
