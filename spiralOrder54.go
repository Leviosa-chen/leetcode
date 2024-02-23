package main

import "fmt"

func main() {
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(nums))
}

func spiralOrder(matrix [][]int) []int {
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1
	var ans []int
	for left <= right && top <= bottom {
		i := left
		for i <= right {
			ans = append(ans, matrix[top][i])
			i++
		}
		i = top + 1
		for i <= bottom {
			ans = append(ans, matrix[i][right])
			i++
		}
		if left < right && top < bottom {
			i = right - 1
			for i > left {
				ans = append(ans, matrix[bottom][i])
				i--
			}
			i = bottom
			for i > top {
				ans = append(ans, matrix[i][left])
				i--
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return ans
}
