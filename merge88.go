package main

import "fmt"

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	l := len(nums1)
	for m > 0 || n > 0 {
		if m > 0 && n > 0 {
			if nums1[m-1] > nums2[n-1] {
				nums1[l-1] = nums1[m-1]
				m--
			} else {
				nums1[l-1] = nums2[n-1]
				n--
			}
		} else if m > 0 {
			nums1[l-1] = nums1[m-1]
			m--
		} else {
			nums1[l-1] = nums2[n-1]
			n--
		}
		l--
	}
}
