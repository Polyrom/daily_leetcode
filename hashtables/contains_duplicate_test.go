/*
217. Contains Duplicate

Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.
*/

package hashtables

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{[]int{1, 2, 4, 2}}, want: true},
		{name: "2", args: args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, want: true},
		{name: "3", args: args{[]int{1, 2, 3, 4}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
