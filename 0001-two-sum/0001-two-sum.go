package leetcode

func twoSum(nums []int, target int) []int {
	for i, ivalue := range nums {
		for j, jvalue := range nums {
			if i != j {
				if ivalue+jvalue == target {
					answer := []int{i, j}
					return answer
				}
			}
		}
	}
	return []int{}
}
