package leetcode

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test case 1", args{"leetcode"}, 0},
		{"Test case 2", args{"loveleetcode"}, 2},
		{"Test case 3", args{"aabb"}, -1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
