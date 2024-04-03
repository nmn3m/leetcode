package leetcode75

func isSubsequence(s string, t string) bool {
	sIndex := 0
	tIndex := 0

	for tIndex < len(t) && sIndex < len(s) {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		tIndex++
	}

	return sIndex == len(s)
}
