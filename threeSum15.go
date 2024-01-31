package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(strs))
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			result = append(result, []int{nums[0], nums[1], nums[2]})
			return result
		}
	}
	sort.Ints(nums)
	start := 0
	for start < len(nums)-2 {
		if nums[start] > 0 {
			return result
		}
		if start > 0 && nums[start] == nums[start-1] {
			start++
			continue
		}
		l := start + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[start] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[start], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
			if sum < 0 {
				l++
			}
			if sum > 0 {
				r--
			}
		}
		start++
	}

	return result
}
