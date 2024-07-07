package main

import "fmt"

func CheckPermutations(str1, str2 string) bool {
	len1, len2 := len(str1), len(str2)
	if len1 != len2 {
		return false
	}

	seen := make(map[rune]int)
	concated := str1 + str2

	for _, c := range concated {
		if _, ok := seen[c]; !ok {
			seen[c] = 1
		} else {
			seen[c] += 1
		}
	}

	for _, c := range seen {
		if c % 2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Check Permutations Challenge")

	str1 := "adcme"
	str2 := "medac"

	isPermutation := CheckPermutations(str1, str2)
	fmt.Println(isPermutation)
}
