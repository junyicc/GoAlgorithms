package dynamic_programming

import "testing"

func TestShortestPathLengthInMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "shortest path length in matrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 9},
					{2, 1, 3, 4},
					{5, 2, 6, 7},
					{6, 8, 4, 3},
				},
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortestPathLengthInMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("ShortestPathLengthInMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
