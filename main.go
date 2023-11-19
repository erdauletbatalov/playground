package main

import (
	"fmt"
)

func main() {
	fmt.Println(asteroidCollision([]int{1, 1, -1, -2}))
}

func asteroidCollision(asteroids []int) []int {
	for i := 0; i < len(asteroids); {
		if i == 0 {
			i++
			continue
		}
		var destroyed int
		if asteroids = destroyRecursively(asteroids, i-1, i, &destroyed); destroyed > 0 {
			i -= destroyed
		} else {
			i++
		}
	}
	return asteroids
}

func destroyRecursively(result []int, leftInd int, rightInd int, destroyed *int) []int {
	if leftInd < 0 || rightInd >= len(result) {
		return result
	}
	if result[leftInd] > 0 && result[rightInd] < 0 {
		if result[leftInd] == -result[rightInd] {
			result = append(result[:leftInd], result[rightInd+1:]...)
			*destroyed++
			return result
		}
		if result[leftInd] > -result[rightInd] {
			result = destroyRecursively(append(result[:rightInd], result[rightInd+1:]...), leftInd, rightInd, destroyed)
			*destroyed++
			return result
		}
		if result[leftInd] < -result[rightInd] {
			result = destroyRecursively(append(result[:leftInd], result[rightInd:]...), leftInd-1, rightInd-1, destroyed)
			*destroyed++
			return result
		}
	}
	return result
}
