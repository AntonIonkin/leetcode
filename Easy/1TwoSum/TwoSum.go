package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	x := make([]int, 0)
	for j := 0; j < len(nums); j++ {
		for i := j + 1; i < len(nums); i++ {
			if nums[j]+nums[i] == target {
				x = append(x, j, i)
				return x
			}
		}
	}
	return x
}
