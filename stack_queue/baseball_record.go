/*
682. Baseball Game

You are keeping the scores for a baseball game with strange rules.
At the beginning of the game, you start with an empty record.

You are given a list of strings operations, where operations[i] is the ith operation you must apply to
the record and is one of the following:

An integer x.
Record a new score of x.
'+'.
Record a new score that is the sum of the previous two scores.
'D'.
Record a new score that is the double of the previous score.
'C'.
Invalidate the previous score, removing it from the record.
Return the sum of all the scores on the record after applying all the operations.
*/

package stack_queue

import "strconv"

type Stack struct {
	vals []int
}

func (s *Stack) Append(v int) {
	s.vals = append(s.vals, v)
}

func (s *Stack) IsEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack) Pop() int {
	last := len(s.vals) - 1
	v := s.vals[last]
	s.vals = s.vals[:last]
	return v
}

func (s *Stack) Peek() int {
	return s.vals[len(s.vals)-1]
}

func (s *Stack) Sum() int {
	var sum int
	for _, v := range s.vals {
		sum += v
	}
	return sum
}

func NewStack(vals []int) *Stack {
	return &Stack{vals: vals}
}

// custom stack solution
// Memory: O(n). Space: O(n)
func calPoints(operations []string) int {
	s := NewStack([]int{})
	for _, op := range operations {
		switch op {
		case "C":
			s.Pop()
		case "D":
			last := s.Peek()
			s.Append(last * 2)
		case "+":
			last := s.Pop()
			prelast := s.Peek()
			s.Append(last)
			s.Append(last + prelast)
		default:
			record, _ := strconv.Atoi(op)
			s.Append(record)
		}
	}
	if s.IsEmpty() {
		return 0
	} else {
		return s.Sum()
	}
}
