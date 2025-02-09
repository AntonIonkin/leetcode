package main

import "fmt"

func main() {
	var str = []string{"flower", "flow", "flight"}
	for _, aa := range str {
		for _, a := range aa {
			fmt.Println(string(a))
		}
	}
}

// func longestCommonPrefix(strs []string) string {

// }
