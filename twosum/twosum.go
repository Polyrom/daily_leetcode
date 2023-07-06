package twosum

import "fmt"

// Return a slice of indices of ints from a slice that add up to target
// At least one correct solution; same element can't be used twice; return in any order

// My solution; time O(n)/mem O(1)
func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}
	for i, j := 0, 1; ; {
		if j == len(nums) && i != len(nums)-2 {
			i++
			j = i + 1
		} else if nums[i]+nums[j] == target {
			return []int{i, j}
		} else {
			j++
		}
	}
}

// Optimized solution; hashmap approach
func twoSumOptimal(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		found, ok := m[n]
		if ok {
			return []int{found, i}
		}
		m[target-n] = i
	}
	return nil
}

func TestTwoSum() {
	testArr1 := []int{11, 15, 2, 7}
	testArr2 := []int{3, 2, 4}
	testArr3 := []int{3, 3}
	testArr4 := []int{3, 2, 3}
	testArr5 := []int{5, 6, 12, 22, 41, 5, 63, 14, 29, 0, 0, 1, 0}
	fmt.Println("My solution")
	fmt.Println("want: [2 3]; have:", twoSum(testArr1, 9))
	fmt.Println("want: [1 2]; have:", twoSum(testArr2, 6))
	fmt.Println("want: [0 1]; have:", twoSum(testArr3, 6))
	fmt.Println("want: [0 2]; have:", twoSum(testArr4, 6))
	fmt.Println("want: [4 6]; have:", twoSum(testArr5, 104))
	fmt.Println("Optimized solution")
	fmt.Println("want: [2 3]; have:", twoSumOptimal(testArr1, 9))
	fmt.Println("want: [1 2]; have:", twoSumOptimal(testArr2, 6))
	fmt.Println("want: [0 1]; have:", twoSumOptimal(testArr3, 6))
	fmt.Println("want: [0 2]; have:", twoSumOptimal(testArr4, 6))
	fmt.Println("want: [4 6]; have:", twoSumOptimal(testArr5, 104))
}
