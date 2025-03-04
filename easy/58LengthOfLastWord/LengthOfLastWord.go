package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := " a"
	s2 := "   fly me   to   the moon  "
	s3 := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s1))
	fmt.Println(lengthOfLastWord(s2))
	fmt.Println(lengthOfLastWord(s3))
}

func lengthOfLastWord(s string) int {
	var start, end int
	n := len(s)

	for i := n - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(s[i])) {
			start = i
			break
		}
	}

	for i := start; i >= 0; i-- {
		if unicode.IsSpace(rune(s[i])) {
			end = i
			break
		}
	}

	if end == 0 {
		return start + 1
	}

	return start - end
}

/*
func lengthOfLastWord(s string) int {
	var start, end int
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(s[i])) {
			start = i
			break
		}
	}
*/
