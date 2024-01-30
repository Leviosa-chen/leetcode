package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char == "M" {
			num += 1000
			continue
		}
		if char == "C" {
			if len(s) > (i+1) && string(s[i+1]) == "D" {
				num += 400
				i++
			} else if len(s) > (i+1) && string(s[i+1]) == "M" {
				num += 900
				i++
			} else {
				num += 100
			}
			continue
		}
		if char == "D" {
			num += 500
			continue
		}

		if char == "X" {
			if len(s) > (i+1) && string(s[i+1]) == "L" {
				num += 40
				i++
			} else if len(s) > (i+1) && string(s[i+1]) == "C" {
				num += 90
				i++
			} else {
				num += 10
			}
			continue
		}
		if char == "L" {
			num += 50
			continue
		}
		if char == "I" {
			if len(s) > (i+1) && string(s[i+1]) == "V" {
				num += 4
				i++
			} else if len(s) > (i+1) && string(s[i+1]) == "X" {
				num += 9
				i++
			} else {
				num += 1
			}
			continue
		}
		if char == "V" {
			num += 5
			continue
		}
	}
	return num
}
