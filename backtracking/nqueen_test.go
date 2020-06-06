package backtracking

import "testing"

func TestCalcNQueen(t *testing.T) {
	type args struct {
		n int
		f func([]int)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "four queen",
			args: args{
				n: 4,
				f: printQueen,
			},
		},
		{
			name: "eight queen",
			args: args{
				n: 8,
				f: printQueen,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CalcNQueen(tt.args.n, tt.args.f)
		})
	}
}
