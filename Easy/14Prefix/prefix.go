package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = []string{"flower", "flight", "flow"}

	fmt.Println(longestCommonPrefix(str))

}

func longestCommonPrefix(strs []string) string {
	prefixSeach := strs[0]

	for _, v := range strs[1:] {

		for !strings.HasPrefix(v, prefixSeach) {

			prefixSeach = prefixSeach[:len(prefixSeach)-1]
			if prefixSeach == "" {
				return ""
			}
		}
	}

	return prefixSeach
}
