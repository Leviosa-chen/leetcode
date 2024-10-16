package main

func main() {

}

func firstMissingPositive(nums []int) int {
	dic := make(map[int]int)
	for _, num := range nums {
		dic[num] = 1
	}

	i := 1
	for ; i <= len(nums); i++ {
		if _, ok := dic[i]; ok {
			continue
		}
		return i
	}
	return i
}
