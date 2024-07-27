"""525. Contiguous Array.

Given a binary array nums, return the maximum length of a contiguous
subarray with an equal number of 0 and 1.
"""


def find_max_length(nums: list[int]) -> int:
    res, count = 0, 0
    sums = {0: -1}
    for i in range(len(nums)):
        count = count + 1 if nums[i] == 1 else count - 1
        if count in sums:
            res = max(res, i - sums[count])
        else:
            sums[count] = i
    return res


print(find_max_length([1, 1, 0, 1, 1, 0, 1, 0]))
