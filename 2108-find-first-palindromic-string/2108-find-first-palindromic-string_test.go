package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Palindrome", "aba", true},
		{"Not Palindrome", "abc", false},
		{"Single Character", "a", true},
		{"Empty String", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isPalindrome(tt.args); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{"First Palindrome Present", []string{"abc", "car", "ada", "racecar", "cool"}, "ada"},
		{"No Palindrome Present", []string{"notapalindrome", "def"}, ""},
		{"Palindrome at the End", []string{"def", "ghi", "racecar"}, "racecar"},
		{"Empty Words", []string{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := firstPalindrome(tt.words); got != tt.want {
				t.Errorf("firstPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
