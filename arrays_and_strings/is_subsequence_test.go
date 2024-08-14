/*
https://leetcode.com/problems/is-subsequence/description/
392. Is Subsequence

Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
*/

package main

import (
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{"abc", "ahbgdc"}, want: true},
		{name: "2", args: args{"axc", "ahbgdc"}, want: false},
		{name: "3", args: args{"ace", "abfegc"}, want: false},
		{name: "4", args: args{"abc", ""}, want: false},
		{name: "5", args: args{"", "abfegc"}, want: true},
		{name: "6", args: args{"", ""}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
