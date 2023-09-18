package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{0}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := range flowerbed {
		if canPlaceFlower(flowerbed, i) {
			n--
			if n == 0 {
				return true
			}
		}
	}
	return false
}

func canPlaceFlower(flowerbed []int, index int) bool {
	if len(flowerbed) == 1 {
		if flowerbed[index] == 0 {
			return true
		} else {
			return false
		}
	}
	switch index {
	case 0:
		if flowerbed[index] == 0 && flowerbed[index+1] == 0 {
			return true
		} else {
			return false
		}
	case len(flowerbed) - 1:
		if flowerbed[index] == 0 && flowerbed[index-1] == 0 {
			return true
		} else {
			return false
		}
	default:
		if flowerbed[index] == 0 && flowerbed[index-1] == 0 && flowerbed[index+1] == 0 {
			return true
		} else {
			return false
		}
	}
}
