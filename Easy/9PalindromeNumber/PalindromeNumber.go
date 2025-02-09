package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 121
	y := -121
	fmt.Println(isPalindrome(x))
	fmt.Println(isPalindrome(y))
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	runes := []rune(str)

	a := 0
	b := len(runes) - 1

	for a < b {
		if runes[a] != runes[b] {
			return false
		}
		a++
		b--
	}
	return true
}
