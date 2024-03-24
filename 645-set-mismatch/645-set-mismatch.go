package leetcode

func findErrorNums(nums []int) []int {

	n := len(nums)
	freq := make(map[int]int)

	missing, duplicate := 0, 0

	for _, num := range nums {
		freq[num]++
	}

	for i := 1; i <= n; i++ {
		if freq[i] == 0 {
			missing = i
		} else if freq[i] == 2 {
			duplicate = i
		}
	}

	return []int{duplicate, missing}
}
