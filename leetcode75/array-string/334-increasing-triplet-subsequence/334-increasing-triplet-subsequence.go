package leetcode75

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	min := nums[0]
	mid := int(^uint(0) >> 1)

	for _, num := range nums {
		if num <= min {
			min = num
		} else if num <= mid {
			mid = num
		} else {
			return true
		}
	}
	return false
}
