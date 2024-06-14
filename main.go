package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("аыфва", "pq"))
}

func mergeAlternately(word1 string, word2 string) string {
	var result []rune
	word1runes := []rune(word1)
	word2runes := []rune(word2)
	if len(word1runes) >= len(word2runes) {
		for i, val := range word2runes {
			result = append(result, word1runes[i], val)
		}
		return string(result) + string(word1runes[len(word2runes):])
	} else {
		for i, val := range word1runes {
			result = append(result, val, rune(word2runes[i]))
		}
		return string(result) + string(word2runes[len(word1runes):])
	}
}
