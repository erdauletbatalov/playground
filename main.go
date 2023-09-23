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

import "fmt"

// initialize map
var vowelsMap = map[rune]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
}

func main() {
	// input string from user stdin
	var input string
	fmt.Scan(&input)

	fmt.Println(reverseVowels(input))
}

func reverseVowels(s string) string {
	var result []rune
	r := findVowels(s)
	for i := range s {
		if isVowel(rune(s[i])) {
			result = append(result, rune(r[len(r)-1]))
			r = r[:len(r)-1]
			continue
		}
		result = append(result, rune(s[i]))
	}
	return string(result)
}

func findVowels(s string) []rune {
	vowels := make([]rune, 0)
	for i := range s {
		if isVowel(rune(s[i])) {
			vowels = append(vowels, rune(s[i]))
		}
	}
	return vowels
}

func isVowel(r rune) bool {
	if _, ok := vowelsMap[r]; ok {
		return true
	}
	return false
}
