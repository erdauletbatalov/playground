package main

import "fmt"

func main() {
	var arr = make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		arr = append(arr, i)
	}
	otherArr := arr
	anotherArr := otherArr
	copy(otherArr[:], otherArr[1:])
	fmt.Println(otherArr)
	anotherArr = append(anotherArr[:], anotherArr[1:]...)
	fmt.Println(anotherArr)
	arr = append(arr[1:], 5)
	fmt.Println(arr)

	copy(otherArr[:], otherArr[1:])
	otherArr = otherArr[:len(otherArr)-1]
}
