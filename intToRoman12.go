package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	var thousands = []string{"", "M", "MM", "MMM"}
	var hundreds = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	var tens = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	var ones = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return thousands[num/1000] + hundreds[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]
}
