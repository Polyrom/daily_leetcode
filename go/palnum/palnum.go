package palnum

import (
	"fmt"
	"strconv"
)

/*
Palindrome Number: given an integer, return true if it's
a palindrome and false otherwise. One-digit number is a palindrome.
*/

// My solution; brute force; conversion to string; time O(n)/mem O(n)
func isPalindrome(x int) bool {
	xs := []byte(strconv.Itoa(x))
	if len(xs) == 1 {
		return true
	}
	var xsRev []byte
	for i := len(xs) - 1; i >= 0; i-- {
		xsRev = append(xsRev, xs[i])
	}
	for j := 0; xs[j] == xsRev[j]; j++ {
		if j == len(xs)-1 {
			return true
		}
	}
	return false
}

// Optimized solution; no conv to string
// Shoutout to https://www.tutorialspoint.com/write-a-golang-program-to-check-whether-a-given-number-is-a-palindrome-or-not
func isPalindromeOptimized(x int) bool {
	if x < 0 {
		return false
	}
	cp := x
	var (
		res       int
		remainder int
	)
	for x > 0 {
		remainder = x % 10
		res = (res * 10) + remainder
		x /= 10
	}
	switch {
	case res == cp:
		return true
	default:
		return false
	}
}

func TestIsPalindrome() {
	var (
		a, b, c, d, e, f, g = 121, 2332, -171, 301203, 23455432, 1000021, 1
	)
	fmt.Println("Palindrome Number. My brute force solution with conversion to string")
	fmt.Println("want: true; have:", isPalindrome(a))
	fmt.Println("want: true; have:", isPalindrome(b))
	fmt.Println("want: false; have:", isPalindrome(c))
	fmt.Println("want: false; have:", isPalindrome(d))
	fmt.Println("want: true; have:", isPalindrome(e))
	fmt.Println("want: false; have:", isPalindrome(f))
	fmt.Println("want: true; have:", isPalindrome(g))
	fmt.Println("Palindrome Number. Optimized solution, no conv to string")
	fmt.Println("want: true; have:", isPalindromeOptimized(a))
	fmt.Println("want: true; have:", isPalindromeOptimized(b))
	fmt.Println("want: false; have:", isPalindromeOptimized(c))
	fmt.Println("want: false; have:", isPalindromeOptimized(d))
	fmt.Println("want: true; have:", isPalindromeOptimized(e))
	fmt.Println("want: false; have:", isPalindromeOptimized(f))
	fmt.Println("want: true; have:", isPalindromeOptimized(g))
}
