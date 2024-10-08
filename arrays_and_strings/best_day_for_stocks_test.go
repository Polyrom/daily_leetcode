package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]int{7, 1, 5, 3, 6, 4}}, want: 5},
		{name: "1", args: args{[]int{2, 2, 5, 3, 6, 12}}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
