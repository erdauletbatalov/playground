package main

import "fmt"

func main() {
	fmt.Println(maxOperations([]int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3))
}

func maxOperations(nums []int, k int) int {
	var result int
	var m = make(map[int]int)
	for i := range nums {
		if nums[i] >= k {
			continue
		}
		if output, ok := m[nums[i]]; ok {
			result++
			if output > 1 {
				m[nums[i]]--
			} else {
				delete(m, nums[i])
			}
			continue
		}
		m[k-nums[i]]++
	}
	return result
}
