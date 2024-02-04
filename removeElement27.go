package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		println(index, nums[i])
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
