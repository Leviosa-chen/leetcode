package main

import "fmt"

func main() {
	arr := []int{0, 0, 2, 8, 9, 8}
	qSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func qSort(list []int, left int, right int) {
	if left >= right {
		return
	}
	index := partition(list, left, right)
	qSort(list, left, index-1)
	qSort(list, index+1, right)
}

func partition(list []int, left int, right int) int {
	target := list[right]
	index := left - 1
	for left < right {
		if list[left] < target {
			index++
			list[left], list[index] = list[index], list[left]
		}
		left++
	}
	index++
	list[right], list[index] = list[index], list[right]
	return index
}
