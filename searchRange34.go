package main

import "fmt"

func main() {
	nums := []int{1}
	fmt.Println(searchRange(nums, 1))
}

func searchRange(nums []int, target int) []int {
	result := findItem(nums, target)
	println(result)
	if result == -1 {
		return []int{-1, -1}
	}
	left := result
	for left >= 0 && nums[left] == target {
		left--
	}
	right := result
	for right < len(nums) && nums[right] == target {
		right++
	}
	return []int{left + 1, right - 1}
}

func findItem(nums []int, target int) int {
	result := -1
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}

	return result
}
