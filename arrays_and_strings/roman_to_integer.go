/*
13. Roman to Integer
https://leetcode.com/problems/roman-to-integer/description/

Given a roman numeral, convert it to an integer.
*/

package main

// Memory: O(n)
// Space: O(1)
func romanToInt(s string) int {
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var sum, greatest int
	for i := len(s) - 1; i >= 0; i-- { // iterate over s right to left
		num := romans[rune(s[i])]
		if num >= greatest { // check if current int is larger than the largest seen previously
			greatest = num // update the greatest number
			sum += num     // add the current int to the total
		} else {
			sum -= num // otherwise, subtract: we hit a subtracting modifier (e.g. IV)
		}
	}
	return sum
}
