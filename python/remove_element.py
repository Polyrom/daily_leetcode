"""
27. Remove element.
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

+ Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
  The remaining elements of nums are not important as well as the size of nums.
+ Return k.
"""


def remove_element(nums: list[int], val: int) -> int:
    if len(nums) == 0:
        return 0
    invalid = -1
    not_equal = 0
    i = 0
    while i <= len(nums) - 1 and nums[i] != invalid:
        if nums[i] == val:
            nums[i] = invalid
            nums.append(nums[i])
            del nums[i]
        else:
            not_equal += 1
            i += 1
    return not_equal


def remove_element_pointers(nums: list[int], val: int) -> int:
    counter = 0
    for i in range(len(nums)):
        if nums[i] != val:
            nums[counter] = nums[i]
            counter += 1
    return counter


def remove_element_fedya(nums: list[int], val: int) -> int:
    val_counter = -1
    for i in range(len(nums) - 1, -1, -1):
        if nums[i] == val:
            nums[i] = nums[val_counter]
            val_counter -= 1
    return len(nums) + val_counter + 1


print(
    "My solution, naive, ugly work with list. R O(n)/ M O(1)."
    "However, not as bad in terms of runtime at LC"
)
nums = [3, 2, 3, 2, 3]
val = 3
print(f"nums: {nums}; val: {val}")
print(f"expected: 2;  actual: {remove_element(nums, val)}")
print(f"nums expected [2,2,_,_]; nums actual: {nums}")

nums = [0, 1, 2, 2, 3, 0, 4, 2]
val = 2
print(f"nums: {nums}; val: {val}")
print(f"expected: 5;  actual: {remove_element(nums, val)}")
print(f"nums expected [0,1,4,0,3,_,_,_]; nums actual: {nums}")

nums = [0, 1, 2, 2, 3, 0, 4, 2]
val = 7
print(f"nums: {nums}; val: {val}")
print(f"expected: 8;  actual: {remove_element(nums, val)}")
print(f"nums expected [0,1,2,2,3,0,4,2]; nums actual: {nums}")


print("Leetcode Solution, two pointers. More elegant, same Big-O, but worse LC runtime")
nums = [3, 2, 3, 2, 3]
val = 3
print(f"nums: {nums}; val: {val}")
print(f"expected: 2;  actual: {remove_element_pointers(nums, val)}")
print(f"nums expected [2,2,_,_]; nums actual: {nums}")

nums = [0, 1, 2, 2, 3, 0, 4, 2]
val = 2
print(f"nums: {nums}; val: {val}")
print(f"expected: 5;  actual: {remove_element_pointers(nums, val)}")
print(f"nums expected [0,1,4,0,3,_,_,_]; nums actual: {nums}")

nums = [0, 1, 2, 2, 3, 0, 4, 2]
val = 7
print(f"nums: {nums}; val: {val}")
print(f"expected: 8;  actual: {remove_element_pointers(nums, val)}")
print(f"nums expected [0,1,2,2,3,0,4,2]; nums actual: {nums}")
