/*
56. Merge Intervals

Given an array of intervals where intervals[i] = [start, end], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/

package main

import (
	"sort"
)

// Memory: O(n log n)
// Space: O(n)
func merge(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool { // sort subslices ascending based on the first element
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 || res[len(res)-1][1] < intervals[i][0] { // if the results is empty OR there is NO overlap
			res = append(res, intervals[i]) // append the current subslice as is
		} else {
			res[len(res)-1][1] = maxInt(res[len(res)-1][1], intervals[i][1]) // if there is an overlap, put maximum value between the current subslice and the current result element
		}
	}
	return res
}
