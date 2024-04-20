"""
169. Majority Element
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ?n / 2? times.
You may assume that the majority element always exists in the array.
"""

from collections import defaultdict


def majority_element(nums: list[int]) -> int:
    m = defaultdict(int)
    res = nums[0]
    max = 0
    for n in nums:
        m[n] += 1
    for k, v in m.items():
        if v > max:
            max = v
            res = k
    return res


def majority_element1(nums: list[int]) -> int:
    return sorted(nums)[len(nums) // 2]


def majority_element2(nums: list[int]) -> int:
    candidate = nums[0]
    count = 0
    for n in nums:
        if count == 0:
            candidate = n
        if n == candidate:
            count += 1
        else:
            count -= 1
    return candidate


print(
    "My solution, primitive. However, beats 70.10% at runtime. Sucks at memory, though."
)
nums = [3, 2, 3]
print(f"nums={nums}; expected: 3; actual: {majority_element(nums)}")
nums = [2, 2, 1, 1, 1, 2, 2]
print(f"nums={nums}; expected: 2; actual: {majority_element(nums)}")

print(
    "My solution, one-liner with sorting. Sucks at runtime and memory, but a one-liner is a one-liner, right?"
)
nums = [3, 2, 3]
print(f"nums={nums}; expected: 3; actual: {majority_element1(nums)}")
nums = [2, 2, 1, 1, 1, 2, 2]
print(f"nums={nums}; expected: 2; actual: {majority_element1(nums)}")

print("Shoutout to Leetcode Solutions, per usual. Moore's Voting Algorithm.")
nums = [3, 2, 3]
print(f"nums={nums}; expected: 3; actual: {majority_element2(nums)}")
nums = [2, 2, 1, 1, 1, 2, 2]
print(f"nums={nums}; expected: 2; actual: {majority_element2(nums)}")
