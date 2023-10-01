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
	input := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println((compress(input)))
}

func compress(chars []byte) int {
	var result []byte
	var charToCompress byte
	var count int = 1
	for i := range chars {
		if len(chars) == 1 {
			result = append(result, chars[i])
			break
		}
		if i == 0 {
			charToCompress = chars[i]
			continue
		}
		if charToCompress == chars[i] {
			count++
			if i == len(chars)-1 {
				if count == 1 {
					result = append(result, charToCompress)
				} else {
					result = append(result, charToCompress)
					result = append(result, intToBytes(count)...)
				}
				break
			}
			continue
		}

		if i == len(chars)-1 {
			if count == 1 {
				result = append(result, charToCompress)
			} else {
				result = append(result, charToCompress)
				result = append(result, intToBytes(count)...)
			}
			charToCompress = chars[i]
			count = 1
		}
		if count == 1 {
			result = append(result, charToCompress)
		} else {
			result = append(result, charToCompress)
			result = append(result, intToBytes(count)...)
		}
		charToCompress = chars[i]
		count = 1
	}
	copy(chars, result)
	fmt.Println(string(chars))
	return len(result)
}

func intToBytes(n int) []byte {
	result := []byte{}
	for n != 0 {
		i := n % 10
		result = append([]byte{byte(i) + '0'}, result...)
		n = n / 10
	}
	return result
}
