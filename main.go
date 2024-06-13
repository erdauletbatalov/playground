package main

import "fmt"

type F interface {
	f()
}

type S1 struct {
	a int
}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

func main() {

	// s1 := []int{1, 2, 3}
	// s2 := s1

	s3 := struct{}{}
	s4 := s3

	s5 := S1{a: 1}
	s6 := S2{}

	fmt.Println(s5 == s6)
	fmt.Println(s3 == s4)
	// fmt.Println(s1 == s2)

	// s1Val := S1{}
	// s1Ptr := &S1{}
	// s2Val := S2{}
	// s2Ptr := &S2{}

	// var i F
	// i = s1Val
	// i = s1Ptr
	// i = s2Val
	// i = s2Ptr

	// fmt.Println(i)

}
