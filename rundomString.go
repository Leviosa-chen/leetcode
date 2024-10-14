package main

import (
	"math/rand"
	"time"
)

func main() {
	println(getRandomString(10, false))
}
func getRandomString(length int, need_special bool) string {
	nomarl_str := "ABCDEF12345abc"
	special_str := "_+-&=!@#$%^*"

	if need_special {
		nomarl_str += special_str
	}
	result := ""
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		randomInt := rand.Intn(len(nomarl_str))
		result += string(nomarl_str[randomInt])
	}
	return result
}
