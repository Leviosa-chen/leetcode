package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

func containsDuplicate(nums []int) bool {
	hasTable := make(map[int]int)
	for _, x := range nums {
		if _, exists := hasTable[x]; exists {
			return true
		} else {
			hasTable[x] = 1
		}
	}
	return false
}
