package leetcode75

func mergeAlternately(word1 string, word2 string) string {
	var merged string
	i, j := 0, 0

	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			merged += string(word1[i])
			i++
		}
		if j < len(word2) {
			merged += string(word2[j])
			j++
		}
	}
	return merged
}
