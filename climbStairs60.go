package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
}

func climbStairs1(n int) int {
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			ans[i] = 1
		} else if i == 1 {
			ans[i] = 2
		} else {
			ans[i] = ans[i-1] + ans[i-2]
		}
	}
	return ans[n-1]
}

func climbStairs(n int) int {
	p := 0
	q := 0
	r := 1

	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
