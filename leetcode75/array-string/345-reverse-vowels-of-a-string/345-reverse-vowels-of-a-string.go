package leetcode75

import "strings"

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"

	sRunes := []rune(s)

	left := 0
	right := len(sRunes) - 1

	for left < right {
		for left < right && !strings.ContainsRune(vowels, sRunes[left]) {
			left++
		}

		for left < right && !strings.ContainsRune(vowels, sRunes[right]) {
			right--
		}

		sRunes[left], sRunes[right] = sRunes[right], sRunes[left]

		left++
		right--

	}

	return string(sRunes)

}
