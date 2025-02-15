package main

import (
	"fmt"
)

func main() {

	input := []string{"h", "e", "l", "l", "o"}

	output := make([]byte, len(input))
	for i, str := range input {
		output[i] = str[0]
	}

	reverseString(output)

	fmt.Println(string(output))

}

func reverseString(s []byte) {
	ln := len(s)

	for i := 0; i < ln/2; i++ {
		s[i], s[ln-1-i] = s[ln-1-i], s[i]
	}

}
