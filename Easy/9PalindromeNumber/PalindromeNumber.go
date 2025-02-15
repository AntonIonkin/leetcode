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

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

/*
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	a := 0
	b := len(str) - 1

	for a < b {
		if str[a] != str[b] {
			return false
		}
		a++
		b--
	}

	return true
}
*/
