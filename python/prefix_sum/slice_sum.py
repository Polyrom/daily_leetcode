"""Sum of elements in a given slice.

Given an array nums, answer multiple queries about the sum of elements
within a specific range [i, j].
"""


def slice_sum(nums: list[int], left: int, right: int) -> int:
    for i in range(len(nums)):
        if i == 0:
            continue
        nums[i] += nums[i - 1]
    return nums[right] - nums[left - 1] if left != 0 else nums[right]
