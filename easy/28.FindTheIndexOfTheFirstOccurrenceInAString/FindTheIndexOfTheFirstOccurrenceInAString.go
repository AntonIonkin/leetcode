package main

import (
	"fmt"
)

func main() {
	haystack := "leetcode"
	needle := "leeto"

	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	index := 0
	j := 0

	for i := 0; i < len(haystack); i++ {
		if haystack[i] != needle[j] {
			index = 0
		}

		if j < len(needle) && haystack[i] == needle[j] {
			j++
		}

		if j == len(needle)-1 {
			return index
		}
	}
	return -1
}

/*
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
*/
