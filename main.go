// Given a string s, reverse only all the vowels in the string and return it.

// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

// Example 1:

// Input: s = "hello"
// Output: "holle"
// Example 2:

// Input: s = "leetcode"
// Output: "leotcede"

// Constraints:

// 1 <= s.length <= 3 * 105
// s consist of printable ASCII characters.

package main

import (
	"fmt"
)

func main() {
	// input string from user stdin
	input := []int{0, 1, 2}
	fmt.Println(productExceptSelf(input))
}

func productExceptSelf(nums []int) []int {

	var result = make([]int, 0, len(nums))
	countOfZero := 0
	for i := range nums {
		result = append(result, 0)
		if nums[i] == 0 {
			countOfZero++
		}
	}
	if countOfZero >= 2 {
		return result
	}
	product := 1
	for i := range result {
		if nums[i] != 0 {
			product *= nums[i]
		}
	}
	if countOfZero == 1 {
		for i := range nums {
			if nums[i] == 0 {
				result[i] = product
				continue
			}
			result[i] = 0
		}
		return result
	}

	for i := range nums {
		result[i] = product / nums[i]
	}
	return result
}
