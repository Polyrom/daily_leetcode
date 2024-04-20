def longest_common_prefix(strs: list[str]) -> str:
    longest_pref = ""
    pref_end = False
    if len(strs) > 0:
        strs.sort()
        first, last = strs[0], strs[len(strs) - 1]
        for i in range(len(first)):
            if not pref_end and first[i] == last[i]:
                longest_pref += first[i]
                continue
            pref_end = True
    return longest_pref


print(longest_common_prefix(["flower", "flight", "flame"]))
print(longest_common_prefix(["hello", "darkness", "my", "old"]))
print(longest_common_prefix(["flower"]))
print(longest_common_prefix([]))
print(longest_common_prefix(["come", "communist", "commonwealth", "comrades"]))
print(longest_common_prefix(["jim", "jimothy", "james", "tuna"]))
print(longest_common_prefix(["ab", "a"]))
print(longest_common_prefix(["abba", "ab", ""]))
print(longest_common_prefix(["hellllo", "hell", "h"]))
