package main

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{[]string{"flower", "flame", "flamingo"}}, want: "fl"},
		{name: "2", args: args{[]string{"groom", "gray", "goose"}}, want: "g"},
		{name: "3", args: args{[]string{"left", "right", "who?"}}, want: ""},
		{name: "4", args: args{[]string{"cir", "car"}}, want: "c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix2(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{[]string{"flower", "flame", "flamingo"}}, want: "fl"},
		{name: "2", args: args{[]string{"groom", "gray", "goose"}}, want: "g"},
		{name: "3", args: args{[]string{"left", "right", "who?"}}, want: ""},
		{name: "4", args: args{[]string{"cir", "car"}}, want: "c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix2(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix2() = %v, want %v", got, tt.want)
			}
		})
	}
}
