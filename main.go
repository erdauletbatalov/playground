package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	var max float64
	var segment float64
	if len(nums) < k {
		return 0
	}
	for i := 0; i < k; i++ {
		segment += float64(nums[i])
	}
	max = segment
	if len(nums) == k {
		return segment / float64(k)
	}
	for i := 1; i < len(nums)-k+1; i++ {
		segment = segment - float64(nums[i-1]) + float64(nums[i+k-1])
		if max < segment {
			max = segment
		}
	}
	return max / float64(k)
}
