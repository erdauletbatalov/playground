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
	input := "the sky is blue"
	fmt.Println(reverseWords(input))
}

func reverseWords(s string) string {
	var result string
	var words []string
	for i := 0; i < len(s); {
		if s[i] != 32 {
			words = append(words, returnWord(s[i:]))
			i = i + len(words[len(words)-1])
			continue
		}
		i++
	}
	for i := len(words) - 1; i >= 0; i-- {
		if i != 0 {
			result += words[i] + " "
			continue
		}
		result += words[i]
	}
	return result
}

func returnWord(s string) string {
	for i := range s {
		if s[i] != 32 {
			continue
		}
		return s[:i]
	}
	return s
}
