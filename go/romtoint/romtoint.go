package romtoint

import "fmt"

/*
Roman To Integer
Given a Roman numeral, return the corresponding integer
*/

// My solution; hard-coded brute force
// time O(n)/ mem O(1)
func romanToInt(s string) int {
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	for i := 0; i < len([]byte(s)); {
		cur := string(s[i])        // current char
		var next string            // next char
		var complement int         // by how many we increment the result
		if i != len([]byte(s))-1 { // find next char if it exists
			next = string(s[i+1])
		}
		if cur == "I" && (next == "V" || next == "X") || // if current char reduces, check following char
			cur == "X" && (next == "L" || next == "C") ||
			cur == "C" && (next == "D" || next == "M") {
			complement = romans[next] - romans[cur]
			i += 2
		} else {
			complement = romans[cur]
			i++
		}
		res += complement
	}
	return res
}

// Optimized solution; same complexity, but more elegant & faster
// Shoutout to https://leetcode.com/problems/roman-to-integer/solutions/187123/20ms-Roman-to-Integer-in-GoLang/
func romanToIntOptimized(s string) int {
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
	for i := len(s) - 1; i >= 0; i-- {
		num := romans[rune(s[i])]
		if num >= greatest {
			greatest = num
			sum += num
		} else {
			sum -= num
		}
	}
	return sum
}

func TestRomanToInt() {
	var (
		a, b, c, d, e = "V", "II", "XIX", "LIV", "CDXLIV"
	)
	fmt.Println("Roman to Integer. My brute force solution")
	fmt.Println("want: 5; have:", romanToInt(a))
	fmt.Println("want: 2; have:", romanToInt(b))
	fmt.Println("want: 19; have:", romanToInt(c))
	fmt.Println("want: 54; have:", romanToInt(d))
	fmt.Println("want: 444; have:", romanToInt(e))
	fmt.Println("Roman to Integer. Optimized solution; more elegant")
	fmt.Println("want: 5; have:", romanToIntOptimized(a))
	fmt.Println("want: 2; have:", romanToIntOptimized(b))
	fmt.Println("want: 19; have:", romanToIntOptimized(c))
	fmt.Println("want: 54; have:", romanToIntOptimized(d))
	fmt.Println("want: 444; have:", romanToIntOptimized(e))
}
