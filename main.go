package main

import (
	"fmt"
)

func main() {
	fmt.Println(gcdOfStrings("LEETLEETLEET", "LEET"))
	fmt.Println(gcdOfStrings("LLLLLLLLLLL", "LLLL"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
}

func gcdOfStrings(str1 string, str2 string) string {
	var x []rune
	var long, short string
	if str1 > str2 {
		long = str1
		short = str2
	} else {
		long = str2
		short = str1
	}
	xMaxLen := findGcd(len(short), len(long))

	for i, val := range short {
		if val != rune(long[i]) {
			return ""
		}
	}

	for i, val := range short {
		if i >= xMaxLen {
			if val == rune(short[i-xMaxLen]) {
				continue
			}
		}

		if val == rune(long[i]) {
			x = append(x, val)
		}
	}
	if len(x) == 0 {
		return ""
	}
	if len(long)%len(x) != 0 {
		return ""
	}
	for i := len(short); i < len(long); i++ {
		if long[i] != long[i-len(short)] {
			return ""
		}
	}

	return string(x)
}

func findGcd(a, b int) int {
	for b != 0 {
		remainder := a % b
		a = b
		b = remainder
	}
	return a
}
