package main

import "fmt"

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums))
}
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
