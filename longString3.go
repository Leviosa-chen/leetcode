package main

import (
	"fmt"
)

func main() {
	str := "abbba"
	fmt.Println(lengthOfLongestSubstring_v2(str))
}

func lengthOfLongestSubstring(s string) int {
	length := 0
	for i := 0; i < len(s); i++ {
		hasTable := make(map[byte]int)
		l := 0
		for j := i; j < len(s); j++ {
			v := s[j]
			if _, exists := hasTable[v]; exists {
				break
			} else {
				l++
				hasTable[v] = i
			}
		}
		if l > length {
			length = l
		}

	}
	return length
}

func lengthOfLongestSubstring_v2(s string) int {
	dict := make(map[byte]int)
	ans := 0
	left := -1
	for right, _ := range s {
		if val, ok := dict[s[right]]; ok {
			// 更新左指针
			if val > left {
				left = val
			}
		}
		// Hash 表中更新为最新的下标
		dict[s[right]] = right
		l := right - left
		if l > ans {
			ans = l
		}
	}
	return ans
}
