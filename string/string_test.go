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
