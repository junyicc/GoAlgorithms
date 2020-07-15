package leetcode

import "testing"

func Test_lengthOfLongestSubstringWithSlidingWin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "length of longest substring",
			args: args{
				s: "ccccabcbcccccc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringWithSlidingWin(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringWithSlidingWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
