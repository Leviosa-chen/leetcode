package main

import "fmt"

func main() {
	nums := []int{2, 6, 7, 8, 11}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hasTable := make(map[int]int)
	for i, x := range nums {
		if value, exists := hasTable[target-x]; exists {
			return []int{i, value}
		} else {
			hasTable[x] = i
		}
	}
	return nil
}
