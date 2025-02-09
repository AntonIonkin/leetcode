package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "rat"
	t := "car"

	fmt.Println(isAnagram(s, t))
}

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
