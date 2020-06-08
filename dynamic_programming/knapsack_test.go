package dynamic_programming

import "testing"

func TestKnapsackWithWeight(t *testing.T) {
	type args struct {
		items []int
		w     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0-1 knapsack with weight",
			args: args{
				items: []int{2, 2, 4, 6, 3},
				w:     9,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KnapsackWithWeight(tt.args.items, tt.args.w); got != tt.want {
				t.Errorf("KnapsackWithWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKnapsackWithValue(t *testing.T) {
	type args struct {
		itemWeight []int
		itemValue  []int
		w          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0-1 knapsack with value",
			args: args{
				itemWeight: []int{2, 2, 4, 6, 3},
				itemValue:  []int{3, 4, 8, 9, 6},
				w:          9,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KnapsackWithValue(tt.args.itemWeight, tt.args.itemValue, tt.args.w); got != tt.want {
				t.Errorf("KnapsackWithValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
