package main

import "fmt"

func main() {
	num1 := "11"
	num2 := "123"
	num3 := "134"

	runeNum1 := []rune(num1)
	runeNum2 := []rune(num2)
	runeNum3 := []rune(num3)

	// lenRuneNum1 := len(runeNum1)
	// lenRuneNum2 := len(runeNum2)

	// if lenRuneNum1 > lenRuneNum2 {
	// 	t := lenRuneNum1 - lenRuneNum2
	// 	result := []rune(num1)[:t]
	// 	for i := lenRuneNum1 - 1; i > t; i-- {
	// 		result = runeNum1[i] - runeNum2[i]
	// 	}
	// }

	fmt.Println(runeNum1)
	fmt.Println(runeNum2)
	fmt.Println(runeNum3)
}
