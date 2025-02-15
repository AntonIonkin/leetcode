package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	min := math.MaxInt64
	max := math.MinInt64
	for _, v := range prices {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}
