package main

func main() {
	strs := []string{"flower", "flow", "flight"}
	println(longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}
	firstStr := strs[0]
outerLoop:
	for i := 0; i < len(firstStr); i++ {
		char := string(firstStr[i])
		for _, str := range strs {
			if i >= len(str) || char != string(str[i]) {
				break outerLoop
			}
		}
		prefix += char
	}

	return prefix
}
