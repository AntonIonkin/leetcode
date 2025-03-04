package main

import "fmt"

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	numbers := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var sum int
	length := len(s)

	for i := 0; i < length; i++ {
		currentValue := numbers[rune(s[i])]
		if i+1 < length {
			nextValue := numbers[rune(s[i+1])]
			if currentValue < nextValue {
				sum -= currentValue
			} else {
				sum += currentValue
			}
		} else {
			sum += currentValue
		}
	}

	return sum
}
