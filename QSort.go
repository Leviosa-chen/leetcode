package main

import "fmt"

func main() {
	arr := []int{5, 7, 5, 3, 1, 8, 6, 4, 2}
	qSort1(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func qSort1(list []int, left int, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	target := list[left]
	for i < j {
		for i < j && list[j] >= target {
			j--
		}
		for i < j && list[i] <= target {
			i++
		}
		if i < j {
			list[i], list[j] = list[j], list[i]
		}
	}
	list[left], list[j] = list[j], list[left]
	qSort1(list, left, j-1)
	qSort1(list, j+1, right)
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
