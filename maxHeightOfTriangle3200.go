package main

func main() {
	println(maxHeightOfTriangle(4, 2))
}

func maxHeightOfTriangle(red int, blue int) int {
	a := getHeight(red, blue)
	b := getHeight(blue, red)
	if a >= b {
		return a
	}
	return b
}

func getHeight(a int, b int) int {
	i := 1
	for true {
		if i%2 == 0 {
			a -= i
		} else {
			b -= i
		}
		if a < 0 || b < 0 {
			break
		}
		i++
	}
	return i - 1
}
