/*
525. Contiguous Array

Given a binary array nums, return the maximum length of a contiguous
subarray with an equal number of 0 and 1.
*/

// Time: O(n) amortized due to hash map look-up.
// Memory: O(n).

/**
 * @param {string[]} nums
 * @return {number}
 */
function FindMaxLength(nums) {
  let count = 0, res = 0;
  const hm = new Map();
  hm.set(count, -1);
  for (let i = 0; i < nums.length; i++) {
    count += nums[i] === 1 ? 1 : -1;
    if (hm.has(count)) {
      res = Math.max(res, i - hm.get(count));
    } else {
      hm.set(count, i);
    }
  }
  return res;
}

module.exports = FindMaxLength;