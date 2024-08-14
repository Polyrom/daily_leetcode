/*
14. Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".
*/

package main

import (
	"math"
	"sort"
	"strings"
)

// Memory: O(n*log n) due to sorting
// Space: O(n)
func longestCommonPrefix(strs []string) string {
	var (
		lcp       strings.Builder
		endPrefix bool
	)
	sort.Strings(strs)
	var first, last string = strs[0], strs[len(strs)-1]
	for i := 0; i < len(first); i++ { // after sorting, we can compare just first and last
		// We can now compare the chars; endPrefix should be evaluated first to avoid IndexError
		if !endPrefix && first[i] == last[i] {
			lcp.WriteByte(first[i])
		} else {
			endPrefix = true // if chars not equal, set the flag, as the prefix has ended
		}
	}
	return lcp.String()
}

// Memory: O((n+k)*n) = O(n), where k = min length of the string
// Space: *O(n) as the hashmap may be of len(n)
func longestCommonPrefix2(strs []string) string {
	var lcp strings.Builder
	minLen := math.MaxInt
	for _, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
		}
	} // find the minimum length
	for i := 0; i < minLen; i++ {
		chars := make(map[byte]interface{}) // initialize a set
		for _, s := range strs {
			chars[s[i]] = nil // add each corresponding element to set
		}
		if len(chars) != 1 {
			return lcp.String() // if there are > 1 elements in the set, return
		}
		lcp.WriteByte(strs[0][i]) // else, write a new element to the string
	}
	return lcp.String()
}
