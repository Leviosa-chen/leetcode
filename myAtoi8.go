package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(myAtoi("  0000000000012345678"))
}

func myAtoi1(s string) int {
	result := 0
	pre := ""
	hasEffect := false
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char == " " && !hasEffect {
			continue
		}
		hasEffect = true
		if pre == "" {
			if char == "-" {
				pre = "-"
				continue
			} else {
				pre = "+"
				if char == "+" {
					continue
				}
			}
		}
		num, err := strconv.Atoi(char)
		if err != nil {
			break
		}
		result = result*10 + num
		resultTmp := result
		if pre == "-" {
			resultTmp = 0 - resultTmp
		}
		if resultTmp < math.MinInt32 {
			return math.MinInt32
		}
		if resultTmp > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	if pre == "-" {
		result = 0 - result
	}
	return result
}

func myAtoi(s string) int {
	// 编译正则表达式
	s = strings.TrimLeftFunc(s, unicode.IsSpace)
	re := regexp.MustCompile(`^[+|\-0-9][0-9]*`)
	// 使用正则表达式进行匹配
	match := re.FindString(s)
	if len(match) == 0 {
		return 0
	}

	num, err := strconv.Atoi(match)
	if err != nil {
		// 防止字符串越界导致的
		if len(match) > 10 {
			if match[0] == '-' {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		return 0
	}

	if num < math.MinInt32 {
		return math.MinInt32
	}
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	// 打印匹配结果
	return num
}
