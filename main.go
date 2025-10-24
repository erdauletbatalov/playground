package main

import "fmt"

func main() {
	arr := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(arr))
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	isInc := true
	var prev int
	for i, num := range arr {
		if i == 0 {
			prev = num
			continue
		}
		if num > prev {
			if !isInc {
				return false
			}
			prev = num
		} else if num == prev {
			return false
		} else if num < prev {
			if i == 1 {
				return false
			}
			isInc = false
			prev = num
		}
	}
	if isInc {
		return false
	}
	return true
}
