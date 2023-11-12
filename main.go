package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3, 3}))
}

func uniqueOccurrences(arr []int) bool {
	firstMap := make(map[int]int)
	secondMap := make(map[int]int)

	for _, val := range arr {
		firstMap[val]++
	}
	for _, val := range firstMap {
		secondMap[val]++
	}
	for _, val := range secondMap {
		if val != 1 {
			return false
		}
	}
	return true
}
