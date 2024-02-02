package main

import "fmt"

func main() {
	fmt.Println(isValid("[]{}"))
}

func isValid(s string) bool {
	var stack []string
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char == "(" || char == "[" || char == "{" {
			stack = append(stack, char)
			continue
		}
		fmt.Println(stack, char, len(stack))
		if len(stack) == 0 {
			return false
		}
		last := len(stack) - 1
		if char == ")" {
			if stack[last] != "(" {
				return false
			}

		}
		if char == "]" {
			if stack[last] != "[" {
				return false
			}
		}
		if char == "}" {
			if stack[last] != "{" {
				return false
			}
		}
		stack = stack[:last]
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
