package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("abcd", "pq"))
}

func mergeAlternately(word1 string, word2 string) string {
	var result []rune
	if len(word1) >= len(word2) {
		for i, val := range word2 {
			result = append(result, rune(word1[i]), val)
		}
		return string(result) + word1[len(word2):]
	} else if len(word1) <= len(word2) {
		for i, val := range word1 {
			result = append(result, val, rune(word2[i]))
		}
		return string(result) + word2[len(word1):]
	}
	return string(result)
}
