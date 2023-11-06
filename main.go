package main

import (
	"fmt"
)

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
func pivotIndex(nums []int) int {
	var sum, left, right int
	for _, v := range nums {
		sum += v
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			left = 0
			right = sum - nums[i]
		} else if i == len(nums)-1 {
			left = sum - nums[i]
			right = 0
		} else {
			left = left + nums[i-1]
			right = right - nums[i]
		}
		if left == right {
			return i
		}
	}
	return -1
}
