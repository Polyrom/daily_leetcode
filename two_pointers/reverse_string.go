/*
344. Reverse String

Write a function that reverses a string. The input string is given as an array of characters s.
You must do this by modifying the input array in-place with O(1) extra memory.
*/

package two_pointers

// Time: O(n). Memory: O(1)
func reverseString(s []byte) {
	var left, right = 0, len(s) - 1
	for left < right {
		var tmp byte
		tmp = s[left]
		s[left] = s[right]
		s[right] = tmp
		left++
		right--
	}
}
