# 88. Merge Sorted Array
# You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, 
# and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

# Merge nums1 and nums2 into a single array sorted in non-decreasing order.

# The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
# To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, 
# and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

def merge(nums1: list[int], m: int, nums2: list[int], n: int) -> None:
    """
    Do not return anything, modify nums1 in-place instead.
    """
    i, j = m - 1, n - 1
    k = m + n - 1
    while j >= 0:
        if i >= 0 and nums1[i] > nums2[j]:
            nums1[k] = nums1[i]
            i -= 1
        else:
            nums1[k] = nums2[j] 
            j -= 1
        k -= 1

print('Optimized O(m+n) solution from Leetcode solutions')
nums1, nums2 = [3, 4, 6, 9, 0, 0, 0, 0], [2, 5, 7, 8]
print(f'nums before: {nums1}; {nums2}')
merge(nums1, 4, nums2, 4)
print(f'nums after expected: [2, 3, 4, 5, 6, 7, 8, 9]; actual: {nums1}')

nums1, nums2 = [1, 2, 3, 0, 0, 0], [4, 5, 6]
print(f'nums before: {nums1}; {nums2}')
merge(nums1, 3, nums2, 3)
print(f'nums after expected: [1, 2, 3, 4, 5, 6]; actual: {nums1}')

nums1, nums2 = [0], [1]
print(f'nums before: {nums1}')
merge(nums1, 0, nums2, 1)
print(f'nums after expected: [1]; actual: {nums1}')