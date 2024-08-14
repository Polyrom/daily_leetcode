/* 2239. Find Closest Number to Zero
https://leetcode.com/problems/find-closest-number-to-zero/description/

Given an integer array nums of size n, return the number with the value closest to 0 in nums.
If there are multiple answers, return the number with the largest value.
*/

package main

// Time: O(n)
// Space: O(1)
func findClosestNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		dist := absInt(nums[i]) // count distance to 0
		if dist < absInt(res) {
			res = nums[i] // if distance is smaller, update res with current value
		} else if dist == absInt(res) {
			res = maxInt(nums[i], res) // if res is equal, update res with larger value
		}
	}
	return res
}
