package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	strs := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(strs, target))
}

func threeSumClosest(nums []int, target int) int {

	result := nums[0] + nums[1] + nums[2]
	if len(nums) == 3 {
		return result
	}
	sort.Ints(nums)
	start := 0
	for start < len(nums)-2 {
		l := start + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[start] + nums[l] + nums[r]
			if math.Abs(float64(sum-target)) < math.Abs(float64(result-target)) {
				result = sum
			}
			if sum > target {
				r--
			}
			if sum < target {
				l++
			}
			if sum == target {
				return result
			}
		}
		start++
	}

	return result
}
