package main

import "fmt"

func main() {
	nums := []int{1, 3, 2}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
