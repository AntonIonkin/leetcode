package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "sadbutsad"
	needle := "sad"

	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
func strStrOther(haystack string, needle string) int {
	index:=0
for i,j:=0,0;i<len(haystack)

	return -1
}
