package longestpref

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
		{"Basic case", args{[]string{"flower", "flight", "flame"}}, "fl"},
		{"No comm pref", args{[]string{"hello", "darkness", "my", "old"}}, ""},
		{"One elem in array", args{[]string{"flower"}}, "flower"},
		{"Empty array", args{[]string{}}, ""},
		{"Basic case #2", args{[]string{"come", "communist", "commonwealth", "comrades"}}, "com"},
		{"No comm pref #2", args{[]string{"jim", "jimothy", "james", "tuna"}}, ""},
		{"Shorter elems later", args{[]string{"ab", "a"}}, "a"},
		{"Empty elem last", args{[]string{"abba", "ab", ""}}, ""},
		{"Shorter elem later #2", args{[]string{"hellllo", "hell", "h"}}, "h"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
