package main

import (
	"fmt"
	"math"
)

func main() {
	x := 8
	fmt.Println(mySqrt(x))
}

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

/*
func mySqrt(x int) int {
	i := 1
	for {
		if i*i > x {
			return i - 1
		}
		i++
	}
}
*/
