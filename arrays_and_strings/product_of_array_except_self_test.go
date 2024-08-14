package main

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{[]int{1, 2, 3, 4}}, want: []int{24, 12, 8, 6}},
		{name: "2", args: args{[]int{-1, 1, 0, -3, 3}}, want: []int{0, 0, 9, 0, 0}},
		{name: "3", args: args{[]int{3, 4, 1, 9, 1}}, want: []int{36, 27, 108, 12, 108}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
