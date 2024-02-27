package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(7, 3))
}

func uniquePaths(m int, n int) int {
	var ans [][]int
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		ans = append(ans, tmp)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 || i == 0 {
				ans[i][j] = 1
			} else {
				ans[i][j] = ans[i-1][j] + ans[i][j-1]
			}
		}
	}

	return ans[n-1][m-1]
}
