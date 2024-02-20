package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	hasTable := make(map[int]int)
	for _, x := range nums {
		if _, exists := hasTable[x]; exists {
			hasTable[x]++
		} else {
			hasTable[x] = 1
		}
	}
	var max, count int
	for i, x := range hasTable {
		fmt.Println(i, x)
		if x > count {
			max = i
			count = x
		}
	}
	return max
}

func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
