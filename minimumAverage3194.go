package main

import (
	"math"
	"sort"
)

func main() {
	println(minimumAverage([]int{7, 8, 3, 4, 15, 13, 4, 1}))
}

func minimumAverage(nums []int) float64 {
	//var averages []float64
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	less_averages := math.MaxFloat64
	for left < right {
		a := float64(nums[left]+nums[right]) / 2
		if less_averages > a {
			less_averages = a
		}
		left++
		right--
	}
	return less_averages
}
