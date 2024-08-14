/*
228. Summary Ranges

You are given a sorted unique integer array nums.
A range [a,b] is the set of all integers from a to b (inclusive).
Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
That is, each element of nums is covered by exactly one of the ranges,
and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:
"a->b" if a != b
"a" if a == b
*/

package main

import (
	"fmt"
)

// Memory: O(n)
// Space: O(n)
func summaryRanges(nums []int) []string {
	var res []string
	curRange := "" // initialize current range
	for i := 0; i < len(nums); i++ {
		// if we encountered the last element or the difference of neighbor elements != 1
		if i == len(nums)-1 || nums[i+1]-nums[i] != 1 {
			curRange += fmt.Sprintf("%d", nums[i]) // add the current element to end of range
			res = append(res, curRange)            // append to the resulting slice
			curRange = ""                          // flush current range
		} else {
			if len(curRange) == 0 { // otherwise, if the current range has not been initialized yet
				curRange += fmt.Sprintf("%d->", nums[i]) // initialize the start of the current range
			}
		}
	}
	return res
}
