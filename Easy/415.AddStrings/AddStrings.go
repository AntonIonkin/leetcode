package main

import "fmt"

func main() {
	num1 := "109"
	num2 := "11"

	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	minNum := []rune(num1)
	maxNum := []rune(num2)

	if len(num1) > len(num2) {
		maxNum, minNum = minNum, maxNum
	}

	j := 0
	remains := 0
	result := []rune{}

	for i := len(maxNum) - 1; i >= 0; i-- {
		sum := int(maxNum[i] - '0')

		if j < len(minNum) {
			sum += int(minNum[len(minNum)-1-j] - '0')
			j++
		}

		sum += remains

		if sum >= 10 {
			remains = 1
			sum -= 10
		} else {
			remains = 0
		}

		result = append([]rune{rune(sum + '0')}, result...)
		fmt.Println(result)
	}

	if remains > 0 {
		result = append([]rune{'1'}, result...)
	}

	return string(result)
}
