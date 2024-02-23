package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(20))
}

func generateMatrix(n int) [][]int {
	left := 0
	right := n - 1
	top := 0
	bottom := n - 1
	var ans [][]int
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		ans = append(ans, tmp)
	}

	index := 1
	for left <= right && top <= bottom {
		i := left
		for i <= right {
			ans[top][i] = index
			index++
			i++
		}
		i = top + 1
		for i <= bottom {
			ans[i][right] = index
			index++
			i++
		}
		if left < right && top < bottom {
			i = right - 1
			for i > left {
				ans[bottom][i] = index
				index++
				i--
			}
			i = bottom
			for i > top {
				ans[i][left] = index
				index++
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
