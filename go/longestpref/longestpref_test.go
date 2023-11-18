package longestpref

import (
	"testing"
)

func TestLongestPref(t *testing.T) {
	a := []string{"flower", "flight", "flame"}
	want := "fl"
	pref := longestCommonPrefixOptimized(a)
	if pref != want {
		t.Fatalf("longestCommonPrefixOptimized = %q, want = %q", pref, want)
	}

}
