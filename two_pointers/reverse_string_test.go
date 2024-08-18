package two_pointers

import (
	"bytes"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"1", args{[]byte("hello")}, []byte("olleh")},
		{"2", args{[]byte("Hannah")}, []byte("hannaH")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !bytes.Equal(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %s, want %s", tt.args.s, tt.want)
			}
		})
	}
}
