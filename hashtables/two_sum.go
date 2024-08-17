/*
1. Two Sum

Given an array of integers nums and an integer target, return indices of the two
numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

package hashtables

func twoSum(nums []int, target int) []int {
	hm := map[int]int{nums[0]: 0}
	for i := 1; i < len(nums); i++ {
		diff := target - nums[i]
		if seen, ok := hm[diff]; ok {
			return []int{seen, i}
		}
		hm[nums[i]] = i
	}
	return []int{}
}
