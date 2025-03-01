package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countAndSay(3))
}

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}

	var result string
	for i := 0; i < n; i++ {
		if i == 0 {
			result = "1"
		}
		if i != 0 {
			result += fmt.Sprintf("%d", strings.Count(result, string(result[i-1])))
		}
	}

	resultRune := []rune(result)
	for i := 0; i < len(result)/2; i++ {
		resultRune[i], resultRune[len(resultRune)-1-i] = resultRune[len(resultRune)-1-i], resultRune[i]
	}

	return string(resultRune)
}
