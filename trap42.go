package main

import "fmt"

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}

func trap1(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		left := 0
		for j := 0; j < i; j++ {
			if height[j] > left {
				left = height[j]
			}
		}
		right := 0
		for j := i + 1; j < len(height); j++ {
			if height[j] > right {
				right = height[j]
			}
		}

		depth := 0
		leftDepth := left - height[i]
		rightDepth := right - height[i]

		if leftDepth > 0 && rightDepth > 0 {
			if leftDepth > rightDepth {
				depth = rightDepth
			} else {
				depth = leftDepth
			}
		}
		println(left, right, leftDepth, rightDepth, depth)
		max += depth

	}
	return max
}

func trap(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	for left < right {
		maxLeft = max(maxLeft, height[left])
		maxRight = max(maxRight, height[right])
		if maxLeft < maxRight {
			ans += maxLeft - height[left]
			left++
		} else {
			ans += maxRight - height[right]
			right--
		}
	}
	return ans
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
