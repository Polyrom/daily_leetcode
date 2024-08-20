/*
167. Two Sum II - Input Array Is Sorted

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
find two numbers such that they add up to a specific target number.
Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.
Your solution must use only constant extra space.
*/

package two_pointers

// Time: O(n). Memory: O(1)
func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)
	var left, right = 0, len(numbers) - 1
	for left <= right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			result[0] = left + 1
			result[1] = right + 1
			return result
		}
		if sum < target || right-left == 1 {
			left++
			right = len(numbers) - 1
			continue
		}
		right--
	}
	return result
}
