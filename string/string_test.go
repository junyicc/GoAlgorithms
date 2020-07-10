package string

import "testing"

func TestBF(t *testing.T) {
	type args struct {
		main    string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "bf1",
			args: args{
				main:    "abcd227fac",
				pattern: "ac",
			},
			want: 8,
		},
		{
			name: "bf2",
			args: args{
				main:    "aaacd227fac",
				pattern: "ac",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BF(tt.args.main, tt.args.pattern); got != tt.want {
				t.Errorf("BF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMP(t *testing.T) {
	type args struct {
		s       string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "kmp1",
			args: args{
				s:       "abc abcdab abcdabcdabde",
				pattern: "bcdabd",
			},
			want: 16,
		},
		{
			name: "kmp2",
			args: args{
				s:       "aabbbbaaabbababbabbbabaaabb",
				pattern: "abab",
			},
			want: 11,
		},
		{
			name: "kmp3",
			args: args{
				s:       "aabbbbaaabbababbabbbabaaabb",
				pattern: "ababacd",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KMP(tt.args.s, tt.args.pattern); got != tt.want {
				t.Errorf("KMP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sort single string",
			args: args{
				s: "dcba",
			},
			want: "abcd",
		}, {
			name: "sort single string",
			args: args{
				s: "dcbbba",
			},
			want: "abbbcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.args.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_match(t *testing.T) {
	type args struct {
		s       string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string regexp0",
			args: args{
				s:       "aab",
				pattern: "c*a*b",
			},
			want: true,
		},
		{
			name: "string regexp1",
			args: args{
				s:       "a",
				pattern: "ab*",
			},
			want: true,
		},
		{
			name: "string regexp2",
			args: args{
				s:       "abcd",
				pattern: ".*..",
			},
			want: true,
		},
		{
			name: "string regexp3",
			args: args{
				s:       "abcd",
				pattern: ".*",
			},
			want: true,
		},
		{
			name: "string regexp4",
			args: args{
				s:       "abcd",
				pattern: ".*c",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match(tt.args.s, tt.args.pattern); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
