"""303. Range Sum Query - Immutable.

Given an integer array nums, handle multiple queries of the following type:

Calculate the sum of the elements of nums between indices left and right
inclusive where left <= right.
Implement the NumArray class:
NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int left, int right) Returns the sum of the elements of nums between
indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
"""


class NumArray:
    """Classic prefix sum approach.

    Create a new array where each element is the sum of previous ones.
    (In-place is not possible due to problem restrictions).
    Then, return array[right] - array[left-1]
    Time: O(n). Memory: O(n)
    """

    def __init__(self, nums: list[int]):
        self.nums = nums

    def sum_range(self, left: int, right: int) -> int:
        nums_summed = []
        for i in range(len(self.nums)):
            if i == 0:
                nums_summed.append(self.nums[i])
                continue
            nums_summed.append(self.nums[i] + nums_summed[i - 1])
        return (
            nums_summed[right] - nums_summed[left - 1]
            if left != 0
            else nums_summed[right]
        )


class NumArray2:
    """Move-pointer approach.

    Move one pointer to the other one and calculate total.
    Time: O(right-left + 1). Memory: O(1).
    """

    def __init__(self, nums: list[int]):
        self.nums = nums

    def sum_range(self, left: int, right: int) -> int:
        total = 0
        while left <= right:
            total += self.nums[left]
            left += 1
        return total
