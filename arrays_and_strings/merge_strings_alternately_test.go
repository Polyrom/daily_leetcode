package main

import "testing"

func Test_mergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{word1: "abc", word2: "pqr"}, want: "apbqcr"},
		{name: "2", args: args{word1: "ab", word2: "pqrf"}, want: "apbqrf"},
		{name: "3", args: args{word1: "abcd", word2: "pq"}, want: "apbqcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternately(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}
