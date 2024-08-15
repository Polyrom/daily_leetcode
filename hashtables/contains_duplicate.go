/*
217. Contains Duplicate

Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.
*/

package hashtables

// Time: O(n). Memory: O(n)
func containsDuplicate(nums []int) bool {
	hm := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := hm[n]; ok {
			return true
		}
		hm[n] = struct{}{}
	}
	return false
}
