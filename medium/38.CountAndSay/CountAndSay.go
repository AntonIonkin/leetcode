package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(3))
}

// 1211
// 1112
func countAndSay(n int) string {
	var l, r int
	base := "1"

	for i := 1; i < n; i++ {
		l, r = 0, 0
		currCount := ""
		for r < len(base) {
			for ; r < len(base) && base[l] == base[r]; r++ {

			}
			currCount += (strconv.Itoa(r-l) + string(base[l]))
			l = r
		}
		base = currCount
	}

	return base
}
