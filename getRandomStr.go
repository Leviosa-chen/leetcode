package main

import (
	"math/rand"
	"time"
)

func main() {
	str := get_randon_str(10, true)
	println(str)
	str = get_randon_str(1000, false)
	println(str)
}

func get_randon_str(l int, need_special_cahr bool) string {
	str_range := "ABCDEFGHKDGabdskdjewekfe1234567890"
	if need_special_cahr {
		str_range = str_range + "!@#$%^&*()"
	}
	result := ""
	rand.Seed(time.Now().UnixNano())
	length := len(str_range)
	for i := 0; i < l; i++ {
		randomInt := rand.Intn(length)
		result += string(str_range[randomInt])
	}
	return result
}
