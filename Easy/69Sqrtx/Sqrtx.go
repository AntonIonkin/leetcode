package main

import "fmt"

func main() {
	x := 9
	fmt.Println(mySqrt(x))
}

func mySqrt(x int) int {
	i := 1
	for {
		if i*i > x {
			return i - 1
		}
		i++
	}
}
