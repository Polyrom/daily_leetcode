package validparanth

import "fmt"

var parantheses = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

// 20.Valid Parantheses
// Determine if input string of parantheses is valid
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// My decision using stack, O(n) time, O(1) memory; beats 63.7% time; 97.6% memory
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	closings := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		_, opening := parantheses[s[i]]
		if opening {
			closings = append(closings, parantheses[s[i]])
		} else {
			if len(closings) == 0 || s[i] != closings[len(closings)-1] {
				return false
			} else {
				closings = closings[:len(closings)-1]
			}
		}
	}
	return len(closings) == 0
}

func TestIsValid() {
	s1, s2, s3, s4, s5, s6 := "()[]{}", "({[]})", "((({}))", "[]{})", "((((({[{{}}]})))))", "(("
	fmt.Println("Valid Parantheses. My solution")
	fmt.Println("want: true; have:", isValid(s1))
	fmt.Println("want: true; have:", isValid(s2))
	fmt.Println("want: false; have:", isValid(s3))
	fmt.Println("want: false; have:", isValid(s4))
	fmt.Println("want: true; have:", isValid(s5))
	fmt.Println("want: false; have:", isValid(s6))
}
