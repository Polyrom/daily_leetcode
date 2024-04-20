package longestpref

import (
	"sort"
)

/*
Longest Common Prefix.
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
