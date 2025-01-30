package main

import "fmt"

func main() {
	fmt.Println(removeStars("st*r**d"))
}

func removeStars(s string) string {
	var result []rune
	for _, val := range s {
		if isChar(val) {
			result = append(result, val)
		} else {
			switch len(result) {
			case 0:
				continue
			default:
				result = result[:len(result)-1]
			}

		}
	}
	return string(result)
}

func isChar(s rune) bool {
	// ASCII alphabetic characters
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}
