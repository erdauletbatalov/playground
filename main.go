package main

import "fmt"

func main() {
	fmt.Println(removeStart("st*r*"))
}

func removeStart(s string) string {
	var result []rune
	for _, val := range s {
		if isChar(val) {
			result = append(result, val)
		} else {
			result = result[:len(result)-1]
		}
	}
	return string(result)
}

func isChar(s rune) bool {
	// ASCII alphabetic characters
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}
