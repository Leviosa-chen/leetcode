package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right])) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func findItem(nums []int, target int) int {
	result := -1
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2
	for mid > left && mid < right {
		println(left, right, mid)
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
