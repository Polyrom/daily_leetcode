def str_str(haystack: str, needle: str) -> int:
    start = -1
    i, j = 0, 0
    while i != len(needle) and j != len(haystack):
        if needle[i] == haystack[j]:
            if i == 0:
                start = j
            i += 1
        else:
            if i != 0:
                j = start
            i = 0
            start = -1
        j += 1
    return start if i == len(needle) else -1


def str_str_optimized(haystack: str, needle: str) -> int:
    for i in range(len(haystack) - len(needle) + 1):
        if haystack[i:i+len(needle)] == needle:
            return i
    return -1


print("My solution, naive, beats almost nothing in terms of memory and space")
haystack = "imsadbutsad"
needle = "sad"
print(f"haystack: {haystack}, needle: {needle}")
print(f"expected: 2, actual: {str_str(haystack, needle)}")

haystack = "sadbutsad"
needle = "sad"
print(f"haystack: {haystack}, needle: {needle}")
print(f"expected: 0, actual: {str_str(haystack, needle)}")

haystack = "mississipi"
needle = "issip"
print(f"haystack: {haystack}, needle: {needle}")
print(f"expected: 4, actual: {str_str(haystack, needle)}")

haystack = "leetcode"
needle = "leeto"
print(f"haystack: {haystack}, needle: {needle}")
print(f"expected: -1, actual: {str_str(haystack, needle)}")

haystack = "aaa"
needle = "aaaa"
print(f"haystack: {haystack}, needle: {needle}")
print(f"expected: -1, actual: {str_str(haystack, needle)}")

print("Shoutout to Leetcode Solutions, optimized")
haystack = "imsadbutsad"
needle = "sad"
print(f"haystack: {haystack}, needle: {needle}")
print(f"expected: 2, actual: {str_str_optimized(haystack, needle)}")

haystack = "sadbutsad"
needle = "sad"
print(f"haystack: {haystack}, needle: {needle}")
print(f"expected: 0, actual: {str_str_optimized(haystack, needle)}")

haystack = "mississipi"
needle = "issip"
print(f"haystack: {haystack}, needle: {needle}")
print(f"expected: 4, actual: {str_str_optimized(haystack, needle)}")

haystack = "leetcode"
needle = "leeto"
print(f"haystack: {haystack}, needle: {needle}")
print(f"expected: -1, actual: {str_str_optimized(haystack, needle)}")

haystack = "aaa"
needle = "aaaa"
print(f"haystack: {haystack}, needle: {needle}")
print(f"expected: -1, actual: {str_str_optimized(haystack, needle)}")
