/*
15. 3Sum

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.
*/

package two_pointers

import "slices"

// Time: O(n2). Memory: O(n) due to sorting.
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)
	for offset := 0; offset < len(nums); offset++ {
		if nums[offset] > 0 {
			break
		}
		if offset > 0 && nums[offset] == nums[offset-1] {
			continue
		}
		low := offset + 1
		high := len(nums) - 1
		for low < high {
			sum := nums[offset] + nums[low] + nums[high]
			if sum == 0 {
				result = append(result, []int{nums[offset], nums[low], nums[high]})
				low++
				high--
				for low < high && nums[low] == nums[low-1] {
					low++
				}
				for low < high && nums[high] == nums[high+1] {
					high--
				}
			} else if sum < 0 {
				low++
			} else {
				high--
			}
		}
	}
	return result
}
