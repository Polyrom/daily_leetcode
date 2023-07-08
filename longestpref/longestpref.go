package longestpref

import (
	"fmt"
	"sort"
)

/*
 Find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string; only lowercase Eng letters
*/

// My solution; brute force; time O(n**2)/ mem O(n); extremely slow & ugly; beats 5% on LeetCode!
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	longComPref := []byte(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < 1 {
			return ""
		}
		for j := 0; j < len(longComPref) && j < len(strs[i]); j++ {
			if []byte(strs[i])[j] != longComPref[j] {
				longComPref = []byte(strs[i])[:j]
			} else if j == len(strs[i])-1 {
				longComPref = []byte(strs[i])
			}
		}
	}
	return string(longComPref)
}

// Optimized solution; time O(n) /mem O(n); Divide and Conquer approach to be added on next review, I guess
// Shoutout to https://medium.com/codex/leetcode-with-golang-longest-common-prefix-89d856b0749b
func longestCommonPrefixOptimized(strs []string) string {
	var (
		longestPrefix string
		endPrefix     bool
	)
	if len(strs) > 0 {
		sort.Strings(strs)
		first, last := string(strs[0]), string(strs[len(strs)-1])
		for i := 0; i < len(first); i++ {
			if !endPrefix && string(first[i]) == string(last[i]) { //!endPrefix should always be evaluated first, IndexError otherwise
				longestPrefix += string(first[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix
}

func TestLongestCommonPrefix() {

	var (
		a = []string{"flower", "flight", "flame"}
		b = []string{"hello", "darkness", "my", "old"}
		c = []string{"flower"}
		d = []string{}
		e = []string{"come", "communist", "commonwealth", "comrades"}
		f = []string{"jim", "jimothy", "james", "tuna"}
		g = []string{"ab", "a"}
		h = []string{"abba", "ab", ""}
	)
	fmt.Println("My solution; brute force; slow & ugly")
	fmt.Println("want: fl; have:", longestCommonPrefix(a))
	fmt.Println("want: ; have:", longestCommonPrefix(b))
	fmt.Println("want: flower; have:", longestCommonPrefix(c))
	fmt.Println("want: ; have:", longestCommonPrefix(d))
	fmt.Println("want: com; have:", longestCommonPrefix(e))
	fmt.Println("want: ; have:", longestCommonPrefix(f))
	fmt.Println("want: a; have:", longestCommonPrefix(g))
	fmt.Println("want: ; have:", longestCommonPrefix(h))
	fmt.Println("Optimized solution; sort and compare first and last strings")
	fmt.Println("want: fl; have:", longestCommonPrefixOptimized(a))
	fmt.Println("want: ; have:", longestCommonPrefixOptimized(b))
	fmt.Println("want: flower; have:", longestCommonPrefixOptimized(c))
	fmt.Println("want: ; have:", longestCommonPrefixOptimized(d))
	fmt.Println("want: com; have:", longestCommonPrefixOptimized(e))
	fmt.Println("want: ; have:", longestCommonPrefixOptimized(f))
	fmt.Println("want: a; have:", longestCommonPrefixOptimized(g))
}
