package leetcode

import "testing"

func Test_defangIPaddr(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test case 1", args{"1.1.1.1"}, "1[.]1[.]1[.]1"},
		{"Test case 2", args{"255.100.50.0"}, "255[.]100[.]50[.]0"},
		{"Test case 3", args{"192.168.0.1"}, "192[.]168[.]0[.]1"},
		{"Test case 4", args{"0.0.0.0"}, "0[.]0[.]0[.]0"},
		{"Test case 5", args{"127.0.0.1"}, "127[.]0[.]0[.]1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defangIPaddr(tt.args.address); got != tt.want {
				t.Errorf("defangIPaddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
