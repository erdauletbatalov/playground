package main

import "fmt"

func main() {
	m := make(map[int]string, 2)
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	appendSlice(arr[1:])
	fmt.Println(arr)

}

func appendSlice(arr []int) {
	arr[0] = 0
	arr[1] = 0

	for i := 0; i < 2; i++ {
		arr = append(arr, i)
	}
	arr[0] = 0
}
