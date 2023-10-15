package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	subs := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] && j+1 < len(t) {
				subs++
				t = t[j+1:]
				break
			} else if s[i] == t[j] {
				subs++
				return subs == len(s)
			}
		}
	}
	return subs == len(s)
}
