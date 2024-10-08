package stack_queue

import "testing"

func Test_isValidParan(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "()[]{}"}, want: true},
		{name: "2", args: args{s: "({[]})"}, want: true},
		{name: "3", args: args{s: "((({}))"}, want: false},
		{name: "4", args: args{s: "[]{})"}, want: false},
		{name: "5", args: args{s: "((((({[{{}}]})))))"}, want: true},
		{name: "6", args: args{s: "(("}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidParan(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
