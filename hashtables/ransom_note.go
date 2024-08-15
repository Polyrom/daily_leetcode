/*
383. Ransom Note

Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using
the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.
*/

package hashtables

// Time: O(n+m). Memory: O(n)
func canConstruct(ransomNote string, magazine string) bool {
	hm := make(map[rune]int)
	for _, char := range magazine {
		if counter, ok := hm[char]; ok {
			hm[char] = counter + 1
			continue
		}
		hm[char] = 1
	}
	foundCounter := 0
	for _, char := range ransomNote {
		if counter, ok := hm[char]; ok {
			hm[char] = counter - 1
			if hm[char] == -1 {
				return false
			}
			foundCounter++
		} else {
			return false
		}
	}
	return len(ransomNote) == foundCounter
}
