package leetcode

func firstUniqChar(s string) int {
	freq := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if freq[s[i]] == 1 {
			return i
		}
	}

	return -1
}
