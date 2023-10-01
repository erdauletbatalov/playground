// 334. Increasing Triplet Subsequence
// Medium
// Topics
// Companies
// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

// Example 1:

// Input: nums = [1,2,3,4,5]
// Output: true
// Explanation: Any triplet where i < j < k is valid.
// Example 2:

// Input: nums = [5,4,3,2,1]
// Output: false
// Explanation: No triplet exists.
// Example 3:

// Input: nums = [2,1,5,0,4,6]
// Output: true
// Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.

// Constraints:

// 1 <= nums.length <= 5 * 105
// -231 <= nums[i] <= 231 - 1

// Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?

package main

import (
	"fmt"
)

func main() {
	// input string from user stdin
	input := []int{1, 2, 3, 4, 5}
	fmt.Println(increasingTriplet(input))
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	difNums := make(map[int]bool)

	for _, val := range nums {
		if _, ok := difNums[val]; !ok {
			difNums[val] = true
		}
	}
	if len(difNums) < 3 {
		return false
	}
	for i := range nums {
		if i+1 >= len(nums) {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j+1 >= len(nums) {
				continue
			}
			if nums[j] > nums[i] {
				for k := j + 1; k < len(nums); k++ {
					if nums[k] > nums[j] {
						return true
					}
				}
			}

		}

	}
	return false
}
