package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello World"
	s2 := "   fly me   to   the moon  "
	s3 := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s1))
	fmt.Println(lengthOfLastWord(s2))
	fmt.Println(lengthOfLastWord(s3))
}

func lengthOfLastWord(s string) int {
	var count int
	s = strings.TrimRight(s, " ")
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			break
		}
		count++
	}
	return count
}
