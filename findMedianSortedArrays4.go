package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, 4}
	nums2 := []int{2, 4, 5, 5}
	fmt.Println(findMedianSortedArrays(nums, nums2))
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	total := l1 + l2
	middleIndex := total / 2
	offset := total % 2
	if offset == 0 {
		middleIndex--
	}
	i := 0
	index1 := 0
	index2 := 0
	current := 0
	for i <= middleIndex {
		if index1 < l1 && index2 < l2 {
			if nums1[index1] > nums2[index2] {
				current = nums2[index2]
				index2++
			} else {
				current = nums1[index1]
				index1++
			}
		} else if index1 >= l1 {
			current = nums2[index2]
			index2++
		} else {
			current = nums1[index1]
			index1++
		}
		i++
	}
	// 单数
	if offset == 1 {
		return float64(current)
	}

	// 双数求下一个
	next := 0
	if index1 < l1 && index2 < l2 {
		if nums1[index1] > nums2[index2] {
			next = nums2[index2]
		} else {
			next = nums1[index1]
		}
	} else if index1 >= l1 {
		next = nums2[index2]
	} else {
		next = nums1[index1]
	}
	return float64((current + next)) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	total := l1 + l2
	middleIndex := total / 2
	offset := total % 2

	// 确保 nums1 是较短的数组
	if l1 > l2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	// 二分法查找中位数
	start := 0
	end := l1
	for start <= end {
		partition1 := (start + end) / 2
		partition2 := middleIndex - partition1

		maxLeft1 := math.MinInt32
		if partition1 > 0 {
			maxLeft1 = nums1[partition1-1]
		}

		minRight1 := math.MaxInt32
		if partition1 < l1 {
			minRight1 = nums1[partition1]
		}

		maxLeft2 := math.MinInt32
		if partition2 > 0 {
			maxLeft2 = nums2[partition2-1]
		}

		minRight2 := math.MaxInt32
		if partition2 < l2 {
			minRight2 = nums2[partition2]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			// 找到了中位数
			if offset == 0 {
				return (math.Max(float64(maxLeft1), float64(maxLeft2)) + math.Min(float64(minRight1), float64(minRight2))) / 2
			} else {
				return math.Min(float64(minRight1), float64(minRight2))
			}
		} else if maxLeft1 > minRight2 {
			// 向左移动 partition1
			end = partition1 - 1
		} else {
			// 向右移动 partition1
			start = partition1 + 1
		}
	}

	return 0.0
}
