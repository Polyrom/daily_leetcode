/*
1768. Merge Strings Alternately
https://leetcode.com/problems/merge-strings-alternately/description/

You are given two strings word1 and word2.
Merge the strings by adding letters in alternating order, starting with word1.
If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.
*/

package main

import (
	"strings"
)

// Time: O(n)
// Space: O(n)
func mergeAlternately(word1 string, word2 string) string {
	var b strings.Builder
	// write elements alternately to the buffer
	// until the end of the shorter string is reached
	for i := 0; i < minInt(len(word1), len(word2)); i++ {
		b.WriteByte(word1[i])
		b.WriteByte(word2[i])
	}
	// write the rest of the longer string to the buffer
	if len(word1) > len(word2) {
		b.WriteString(word1[len(word2):])
	}
	if len(word2) > len(word1) {
		b.WriteString(word2[len(word1):])
	}
	return b.String()
}
