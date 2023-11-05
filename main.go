package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
}

func largestAltitude(gain []int) int {
	var altitude []int
	for i := 0; i < len(gain)+1; i++ {
		if i == 0 {
			altitude = append(altitude, 0)
			continue
		}
		altitude = append(altitude, altitude[i-1]+gain[i-1])
	}
	return max(altitude)
}

func max(arr []int) int {
	var max int
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
