package main

import "fmt"

func main() {
	fmt.Println(1 << 3)
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
