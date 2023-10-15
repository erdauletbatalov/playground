package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 0, 0, 1})
	moveZeroes([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0})
}

func moveZeroes(nums []int) {
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j-1] == 0 {
					zeroCount++
				}
				if nums[j] != 0 {
					nums[j-zeroCount], nums[j] = nums[j], nums[j-zeroCount]
					zeroCount--
				}
			}
			break
		}
	}
	fmt.Println(nums)
}
