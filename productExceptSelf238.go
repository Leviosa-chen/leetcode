package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ansL := make([]int, n)
	ansR := make([]int, n)
	ans := make([]int, n)
	ansL[0] = 1
	for i := 1; i < n; i++ {
		ansL[i] = ansL[i-1] * nums[i-1]
	}
	ansR[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		ansR[i] = ansR[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		ans[i] = ansL[i] * ansR[i]
	}
	return ans
}
