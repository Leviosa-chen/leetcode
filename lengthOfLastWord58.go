package main

import "strings"

func main() {

}

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	arr := strings.Split(s, " ")
	return len(arr[len(arr)-1])
}
