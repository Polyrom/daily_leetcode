"""
26. Remove Duplicates from Sorted Array
Given an integer array nums sorted in non-decreasing order, remove the duplicates
in-place such that each unique element appears only once.
The relative order of the elements should be kept the same.
Then return the number of unique elements in nums.
"""


def remove_duplicates(nums: list[int]) -> int:
    if len(nums) <= 1:
        return len(nums)

    counter = 1
    i = 0
    while i <= len(nums) - 2 and nums[i] <= nums[i + 1]:
        if nums[i] == nums[i + 1]:
            del nums[i + 1]
        else:
            counter += 1
            i += 1
    return counter


def remove_duplicates_optimized(nums: list[int]) -> int:
    slow = 0
    for fast in range(1, len(nums)):
        if nums[slow] != nums[fast]:
            nums[slow + 1] = nums[fast]
            slow += 1
    return slow + 1


def remove_duplicates_optimized_again(nums: list[int]) -> int:
    counter = 0
    for i in range(len(nums) - 1):
        if nums[i] != nums[i + 1]:
            nums[counter] = nums[i]
            counter += 1
    if nums[len(nums) - 1] != nums[counter]:
        nums[counter] = nums[len(nums) - 1]
        counter += 1
    if nums[counter - 1] != nums[len(nums) - 1]:
        counter += 1
    return max(counter, 1)


print("my solution")
a = [1, 1, 2]
print(
    f"expected: 2, actual: {remove_duplicates(a)}\n"
    f"nums before:[1, 1, 2], after: {a}"
)

a = [1, 1, 1, 2, 2, 4, 5, 5, 5, 6]
print(
    f"expected: 5, actual: {remove_duplicates(a)}\n"
    f"nums before:[1, 1, 2], after: {a}"
)

a = [1, 1, 1]
print(
    f"expected: 1, actual: {remove_duplicates(a)}\n"
    f"nums before:[1, 1, 1], after: {a}"
)

a = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
print(
    f"expected: 5, actual: {remove_duplicates(a)}\n"
    f"nums before:[0,0,1,1,1,2,2,3,3,4], after: {a}"
)


print("optimized solution")
a = [1, 1, 2]
print(
    f"expected: 2, actual: {remove_duplicates_optimized(a)}\n"
    f"nums before:[1, 1, 2], after: {a}"
)

a = [1, 1, 1, 2, 2, 4, 5, 5, 5, 6]
print(
    f"expected: 5, actual: {remove_duplicates_optimized(a)}\n"
    f"nums before:[1, 1, 2], after: {a}"
)

a = [1, 1, 1]
print(
    f"expected: 1, actual: {remove_duplicates_optimized(a)}\n"
    f"nums before:[1, 1, 1], after: {a}"
)

a = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
print(
    f"expected: 5, actual: {remove_duplicates_optimized(a)}\n"
    f"nums before:[0,0,1,1,1,2,2,3,3,4], after: {a}"
)

print("optimized solution, better but more cumbersome")
a = [1, 1, 2]
print(
    f"expected: 2, actual: {remove_duplicates_optimized_again(a)}\n"
    f"nums before:[1, 1, 2], after: {a}"
)

a = [1, 1, 1, 2, 2, 4, 5, 5, 5, 6]
print(
    f"expected: 5, actual: {remove_duplicates_optimized_again(a)}\n"
    f"nums before:[1, 1, 2], after: {a}"
)

a = [1, 1, 1]
print(
    f"expected: 1, actual: {remove_duplicates_optimized_again(a)}\n"
    f"nums before:[1, 1, 1], after: {a}"
)

a = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
print(
    f"expected: 5, actual: {remove_duplicates_optimized_again(a)}\n"
    f"nums before:[0,0,1,1,1,2,2,3,3,4], after: {a}"
)
