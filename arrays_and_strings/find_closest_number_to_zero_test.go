package main

import "testing"

func Test_findClosestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{nums: []int{-4, -2, 1, 4, 8}}, want: 1},
		{name: "2", args: args{nums: []int{-4, -2, 2, 4, 8}}, want: 2},
		{name: "3", args: args{nums: []int{-4, -2, 0, 2, 4, 8}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestNumber(tt.args.nums); got != tt.want {
				t.Errorf("findClosestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
