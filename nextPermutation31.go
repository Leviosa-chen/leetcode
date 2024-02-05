package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverseNums(nums[i+1:])
}

func reverseNums(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}
