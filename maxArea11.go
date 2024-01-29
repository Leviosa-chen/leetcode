package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea1(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			h := height[i]
			if height[i] > height[j] {
				h = height[j]
			}
			area := (j - i) * h
			if max < area {
				max = area
			}
		}
	}
	return max
}
