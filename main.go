package main

import "fmt"

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))
}
func replaceElements(arr []int) []int {
	var max int
	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			max = arr[i]
			arr[i] = -1
			continue
		}
		// if len(arr) != 2 && i == len(arr)-2 {
		// 	temp = max
		// 	if arr[i] > max {
		// 		max = arr[i]
		// 	}
		// 	arr[i] = temp
		// 	continue
		// }
		temp := max
		if arr[i] > max {
			max = arr[i]
		}
		arr[i] = temp
	}
	return arr
}
