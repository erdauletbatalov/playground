package main

import "fmt"

func main() {
	ptr := new(bool)
	var a *bool
	var b *bool
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println()
	fmt.Println(a)
	a = new(bool)
	fmt.Println(*a)
	fmt.Println(*b)
}
