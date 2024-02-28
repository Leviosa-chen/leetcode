package main

import "fmt"

func main() {
	nums := []int{1, 3, 2}
	fmt.Println(subsets(nums))
}
func subsets(nums []int) [][]int {
	n := len(nums)
	var dfs func(int, []int)
	var ans [][]int
	dfs = func(i int, tmp []int) {
		ans = append(ans, append([]int(nil), tmp...))
		for j := i; j < n; j++ {
			dfs(j+1, append(tmp, nums[j]))
		}
	}
	dfs(0, []int{})
	return ans
}
