/*
125. Valid Palindrome

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing
all non-alphanumeric characters, it reads the same forward and backward.
Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/

package two_pointers

import (
	"strings"
	"unicode"
)

func isAlphaNumeric(char byte) bool {
	return unicode.IsLetter(rune(char)) || unicode.IsDigit(rune(char))
}

// Time: O(n). Memory: O(1)
func isPalindrome(s string) bool {
	var left, right = 0, len(s) - 1
	for left < right {
		if !isAlphaNumeric(s[left]) {
			left++
			continue
		}
		if !isAlphaNumeric(s[right]) {
			right--
			continue
		}
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}
