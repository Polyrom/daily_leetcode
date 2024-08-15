/*
771. Jewels and Stones

You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have.
Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.

Letters are case sensitive, so "a" is considered a different type of stone from "A".
*/

package hashtables

// Time: O(n+m), Memory: O(n)
func numJewelsInStones(jewels string, stones string) int {
	var count int
	hs := make(map[rune]struct{})
	for _, j := range jewels {
		hs[j] = struct{}{}
	}
	for _, s := range stones {
		if _, ok := hs[s]; ok {
			count++
		}
	}
	return count
}
