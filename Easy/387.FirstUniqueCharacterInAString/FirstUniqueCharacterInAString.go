package main

import "fmt"

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	var index [26]int

	for _, v := range s {
		index[v-97]++
	}

	for i, v := range s {
		if index[v-97] == 1 {
			return i
		}
	}

	return -1
}

/*
func firstUniqChar(s string) int {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	for i, char := range s {
		if charCount[char] == 1 {
			return i
		}
	}

	return -1
}
*/
