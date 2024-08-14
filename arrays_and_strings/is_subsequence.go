/*
https://leetcode.com/problems/is-subsequence/description/
392. Is Subsequence

Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
*/

package main

// Time: O(n)
// Space: O(1)
func isSubsequence(s string, t string) bool {
	var si, ti int = 0, 0            // initialize pointers for each string
	for si < len(s) && ti < len(t) { // iterate over both strings until we reach the last char
		if s[si] == t[ti] {
			si++ // if chars are equal, move the substring pointer right
		}
		ti++ // move string-under-search pointer right on each iteration
	}
	return si == len(s) // if we moved the searched string pointer to the end, the substring is found
}
