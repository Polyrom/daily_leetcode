def two_sum(nums: list[int], target: int) -> list[int]:
    m = {}
    for i, n in enumerate(nums):
        if m.get(n) is not None:
            return [i, m[n]]
        m[target - n] = i
    return []


print(two_sum([11, 15, 2, 7], 9))
print(two_sum([3, 2, 3], 6))
print(two_sum([5, 6, 12, 22, 41, 5, 63, 14, 29, 0, 0, 1, 0], 104))
