package main

import "fmt"

func main() {
	fmt.Println(closeStrings("cabbba", "abbccc"))
}

func closeStrings(word1 string, word2 string) bool {
	firstMap := make(map[rune]int)
	var firstSum int
	secondMap := make(map[rune]int)
	var secondSum int
	thirdMap := make(map[int]int)
	forthMap := make(map[int]int)

	for _, val := range word1 {
		firstMap[val]++
	}
	for _, val := range word2 {
		secondMap[val]++
	}
	for key, val := range firstMap {
		if _, ok := secondMap[key]; !ok {
			return false
		}
		firstSum += val
	}
	for key, val := range secondMap {
		if _, ok := firstMap[key]; !ok {
			return false
		}
		secondSum += val
	}
	if firstSum != secondSum {
		return false
	}
	for _, val := range firstMap {
		thirdMap[val]++
	}
	for _, val := range secondMap {
		forthMap[val]++
	}
	if len(firstMap) != len(secondMap) {
		return false
	}
	for thirdKey, thirdVal := range thirdMap {
		if forthVal, ok := forthMap[thirdKey]; ok && thirdVal == forthVal {
			continue
		}
		return false
	}
	return true
}
