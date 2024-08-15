/*
242. Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
*/

package hashtables

// Time: O(n + m). Memory: O(n)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hm := make(map[rune]int)
	for _, c := range s {
		hm[c]++
	}
	for _, c := range t {
		if _, ok := hm[c]; ok {
			hm[c]--
		}
		if hm[c] == 0 {
			delete(hm, c)
		}
	}
	return len(hm) == 0
}
