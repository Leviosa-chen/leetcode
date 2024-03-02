package main

import "fmt"

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit1(nums))
}
func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	min := prices[0]
	for _, v := range prices {
		if v-min > max {
			max = v - min
		}
		if v < min {
			min = v
		}
	}
	return max
}
