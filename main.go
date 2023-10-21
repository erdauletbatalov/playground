package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	var l, result int
	r := len(height) - 1

	for l < r {
		currArea := min(height[l], height[r]) * (r - l)
		result = max(result, currArea)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Brute force
// func maxArea(height []int) int {
// 	max := 0
// 	for i := 0; i < len(height); i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			if height[i] >= height[j] {
// 				area := height[j] * (j - i)
// 				if area > max {
// 					max = area
// 				}
// 			} else if height[i] < height[j] {
// 				area := height[i] * (j - i)
// 				if area > max {
// 					max = area
// 				}
// 			}
// 		}
// 	}
// 	return max
// }
