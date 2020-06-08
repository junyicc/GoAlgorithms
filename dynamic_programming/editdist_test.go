package dynamic_programming

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lcs",
			args: args{
				s1: "abcd",
				s2: "aacd",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubsequence(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("LongestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
