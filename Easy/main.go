package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "[Лия][Парфенова]"
	strSlice := strings.Split(str, "][")

	for i, v := range strSlice {
		fmt.Printf("%d-%s\n", i, v)
	}
}
