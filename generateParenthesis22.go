package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(8))
}

func generateParenthesis1(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	data := generateParenthesis(n - 1)
	hasTable := make(map[string]int)
	var result []string
	for _, str := range data {
		l := len(str)
		tmp := ""
		for i := 0; i < l; i++ {
			tmp = str[0:i] + "()" + str[i:l]
			if _, exists := hasTable[tmp]; exists {
				continue
			} else {
				hasTable[tmp] = i
				result = append(result, tmp)
			}
		}

		tmp = "(" + str + ")"
		if _, exists := hasTable[tmp]; exists {
			continue
		} else {
			hasTable[tmp] = 1
			result = append(result, tmp)
		}
	}
	return result
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	dp := make([][]string, n+1)
	dp[0] = []string{""}

	for i := 1; i <= n; i++ {
		combinations := []string{}
		for j := 0; j < i; j++ {
			for _, str1 := range dp[j] {
				for _, str2 := range dp[i-j-1] {
					combinations = append(combinations, "("+str1+")"+str2)
				}
			}
		}
		dp[i] = combinations
	}

	return dp[n]
}
