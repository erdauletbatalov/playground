package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 1}))
}

func longestSubarray(nums []int) int {
	var prev, zeros, ones, max, curr int

	for i := range nums {
		if nums[i] == 1 {
			ones++
		} else {
			zeros++
		}
	}
	if ones == len(nums) {
		return len(nums) - 1
	}
	if zeros == len(nums) {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			curr++
		} else if nums[i] == 0 && i != 0 && i != len(nums)-1 {
			if nums[i+1] == 0 {
				prev = 0
				if max < prev+curr {
					max = prev + curr
				}
				curr = 0
			} else {
				if max < prev+curr {
					max = prev + curr
				}
				prev = curr
				curr = 0
			}
		}
		if max < prev+curr {
			max = prev + curr
		}
	}
	return max
}
