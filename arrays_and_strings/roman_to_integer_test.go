/*
13. Roman to Integer
Given a roman numeral, convert it to an integer.
*/

package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{"V"}, want: 5},
		{name: "1", args: args{"II"}, want: 2},
		{name: "1", args: args{"XIX"}, want: 19},
		{name: "1", args: args{"LIV"}, want: 54},
		{name: "1", args: args{"CDXLIV"}, want: 444},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
