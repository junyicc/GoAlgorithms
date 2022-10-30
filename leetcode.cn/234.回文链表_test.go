package leetcodecn

import "testing"

func Test_isPalindromeByReverse(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{head: makeList([]int{1, 2, 2, 1})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeByReverse(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeByReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
