package main

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{[]int{0, 1, 2, 4, 5, 7}}, want: []string{"0->2", "4->5", "7"}},
		{name: "2", args: args{[]int{0, 2, 3, 4, 6, 8, 9}}, want: []string{"0", "2->4", "6", "8->9"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
