/*
238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i]
is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit
in a 32-bit integer.

You must write an algorithm that runs in O(n) time
and without using the division operation.
*/

package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))            // initialize the final res of len(nums)
	var leftProduct, rightProduct int = 1, 1 // initialize products
	for i := 0; i < len(nums); i++ {
		res[i] = leftProduct
		leftProduct *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return res
}
