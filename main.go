package main

import "fmt"

func main() {
	arr := []int{10, 2, 5, 3}
	fmt.Println(checkIfExist(arr))

	arr = []int{3, 1, 7, 11}
	fmt.Println(checkIfExist(arr))

}
func checkIfExist(arr []int) bool {
	seen := make(map[int]bool)
	for _, num := range arr {
		// Check if 2 * num exists in the set
		if seen[num*2] {
			return true
		}
		// Check if num / 2 exists in the set (only if num is even)
		if num%2 == 0 && seen[num/2] {
			return true
		}
		// Add current number to the set
		seen[num] = true
	}
	return false
}
