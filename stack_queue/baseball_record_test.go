package stack_queue

import "testing"

func Test_calPoints(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"5", "2", "C", "D", "+"}}, 30},
		{"2", args{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}}, 27},
		{"3", args{[]string{"1", "C"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calPoints(tt.args.operations); got != tt.want {
				t.Errorf("calPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
