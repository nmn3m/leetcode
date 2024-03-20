package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", struct {
			nums   []int
			target int
		}{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"Example 2", struct {
			nums   []int
			target int
		}{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"Example 3", struct {
			nums   []int
			target int
		}{[]int{3, 3}, 6}, []int{0, 1}},
		{"No solution", struct {
			nums   []int
			target int
		}{[]int{1, 2, 3, 4}, 10}, []int{}},
		{"Empty input", struct {
			nums   []int
			target int
		}{[]int{}, 10}, []int{}},
		{"Single element", struct {
			nums   []int
			target int
		}{[]int{5}, 5}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
