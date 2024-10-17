package main

func main() {
	println(mySqrt(0))
}

func mySqrt(x int) int {
	left := 0
	right := x
	ret := -1
	for left <= right {
		mid := (left + right) / 2
		tmp := mid * mid
		if tmp == x {
			return mid
		} else if tmp < x {
			ret = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ret
}
