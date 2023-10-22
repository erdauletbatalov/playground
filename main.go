package main

import "fmt"

func main() {
	fmt.Println(maxVowels("hellooo", 4))
}

func maxVowels(s string, k int) int {
	curr := 0
	vowels := "aeiou"
	var max int
	if len(s) < k {
		return 0
	}
	for i := 0; i < k; i++ {
		if contains(vowels, s[i]) {
			max++
		}
		curr = max
	}
	for i := 1; i < len(s)-k+1; i++ {
		if contains(vowels, s[i-1]) {
			curr--
		}
		if contains(vowels, s[i+k-1]) {
			curr++
		}
		if curr > max {
			max = curr
		}
	}
	return max
}

func contains(str string, ch byte) bool {
	for i := range str {
		if ch == str[i] {
			return true
		}
	}
	return false
}
