package main

import "fmt"

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 2}))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var hashMap map[int]int
	hashMap = make(map[int]int)
	firstMap := make(map[int]bool)
	secondMap := make(map[int]bool)
	var result = [][]int{[]int{}, []int{}}

	for _, val := range nums1 {
		hashMap[val] = -1
	}
	for _, val := range nums2 {
		if res, ok := hashMap[val]; !ok {
			hashMap[val] = 1
			continue
		} else if ok && res != 1 {
			hashMap[val] = 0
		}
	}
	for _, val := range nums1 {
		if res := hashMap[val]; res == -1 {
			if _, ok := firstMap[val]; !ok {
				firstMap[val] = true
				result[0] = append(result[0], val)
			}

		}
	}
	for _, val := range nums2 {
		if res := hashMap[val]; res == 1 {
			if _, ok := secondMap[val]; !ok {
				secondMap[val] = true
				result[1] = append(result[1], val)
			}
		}
	}
	return result
}
