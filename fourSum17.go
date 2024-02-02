package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{-3, -1, 0, 2, 4, 5}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {

	result := [][]int{}
	if len(nums) < 4 {
		return result
	}
	if len(nums) == 4 {
		if nums[0]+nums[1]+nums[2]+nums[3] == target {
			result = append(result, []int{nums[0], nums[1], nums[2], nums[3]})
			return result
		}
	}
	sort.Ints(nums)
	start := 0
	n := len(nums)
	for start < n-3 {
		if nums[start]+nums[start+1]+nums[start+2]+nums[start+3] > target {
			return result
		}
		if start > 0 && nums[start] == nums[start-1] || nums[start]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			start++
			continue
		}
		index := start + 1
		for index < n-2 {
			if index > start+1 && nums[index] == nums[index-1] || nums[start]+nums[index]+nums[n-2]+nums[n-1] < target {
				index++
				continue
			}
			l := index + 1
			r := len(nums) - 1
			for l < r {
				sum := nums[start] + nums[index] + nums[l] + nums[r]
				if sum == target {
					result = append(result, []int{nums[start], nums[index], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				}
				if sum < target {
					l++
				}
				if sum > target {
					r--
				}
			}
			index++
		}

		start++
	}

	return result
}
