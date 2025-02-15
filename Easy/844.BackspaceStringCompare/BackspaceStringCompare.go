package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "xywrrmp"
	t := "xywrrmu#p"

	fmt.Println(backspaceCompare(s, t))

}

func backspaceCompare(s string, t string) bool {
	sRune := []rune(s)
	tRune := []rune(t)

	count := 0
	for i := len(sRune) - 1; i >= 0; i-- {
		if sRune[i] == '#' {
			sRune[i] = '_'
			count++
		}
		if unicode.IsLetter(sRune[i]) && count != 0 {
			sRune[i] = '_'
			count--
		}
	}

	count = 0
	for i := len(tRune) - 1; i >= 0; i-- {
		if tRune[i] == '#' {
			tRune[i] = '_'
			count++
		}
		if unicode.IsLetter(tRune[i]) && count != 0 {
			tRune[i] = '_'
			count--
		}
	}

	s = strings.ReplaceAll(string(sRune), "_", "")
	t = strings.ReplaceAll(string(tRune), "_", "")

	return s == t
}
