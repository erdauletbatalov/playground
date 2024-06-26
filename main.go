package main

import "fmt"

type Foo struct {
	a string
}

func (f Foo) A(s string) string {
	return s
}
func (f Foo) B(s string) string {
	return s
}
func (f Foo) C(s string) string {
	return s
}

type AB interface {
	A(s string) string
	B(s string) string
}

type BC interface {
	B(s string) string
	C(s string) string
}

func main() {
	var f AB = Foo{}
	fmt.Println(f.B("ooh hell"))
	newF := f.(BC)
	fmt.Println(newF.A("naah"))
}
