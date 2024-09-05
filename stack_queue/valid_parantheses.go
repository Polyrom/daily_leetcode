package stack_queue

func isValidParan(s string) bool {
	ps := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := make([]string, 0)

	for _, p := range s {
		if paran, isClosing := ps[string(p)]; isClosing {
			if len(stack) == 0 {
				return false
			}
			lastInStack := len(stack) - 1
			opening := stack[lastInStack]
			if paran != opening {
				return false
			}
			stack = stack[:lastInStack]
		} else {
			stack = append(stack, string(p))
		}
	}
	return len(stack) == 0
}
