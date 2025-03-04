package main

import (
	"fmt"
)

func main() {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	t := "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbba"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	var sByte [26]int
	var tByte [26]int

	for _, v := range s {
		sByte[v-97]++
	}

	for _, v := range t {
		tByte[v-97]++
	}

	fmt.Println(sByte) //[1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(tByte) //[1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	return sByte == tByte
}

/*
func isAnagram(s string, t string) bool {
	runeS := []rune(s)
	runeT := []rune(t)

	sort.Slice(runeS, func(i, j int) bool {
		return runeS[i] < runeS[j]
	})

	sort.Slice(runeT, func(i, j int) bool {
		return runeT[i] < runeT[j]
	})

	s = string(runeS)
	t = string(runeT)
	return s == t
}
*/
