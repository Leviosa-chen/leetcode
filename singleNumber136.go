package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

func singleNumber1(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	for i := 0; i < l-1; i++ {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
		i++
	}
	return nums[l-1]
}

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
