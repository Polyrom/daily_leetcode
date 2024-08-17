/*
977. Squares of a Sorted Array
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each
number sorted in non-decreasing order.
*/

package two_pointers

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	var left, right, cur = 0, len(nums) - 1, len(nums) - 1
	for left <= right {
		var powLeft, powRight = nums[left] * nums[left], nums[right] * nums[right]
		if powLeft >= powRight {
			result[cur] = powLeft
			left++
		} else {
			result[cur] = powRight
			right--
		}
		cur--
	}
	return result
}
