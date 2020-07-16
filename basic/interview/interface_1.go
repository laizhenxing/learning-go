package main

import "fmt"

type S3 struct{}

func (s1 S3) f() {
	fmt.Println("S1.f()")
}
func (s1 S3) g() {
	fmt.Println("S1.g()")
}

type S2 struct {
	S3
}

func (s2 S2) f() {
	fmt.Println("S2.f()")
}

type I interface {
	f()
}

func printType(i I) {

	fmt.Printf("%T\n", i)
	if s1, ok := i.(S3); ok {
		s1.f()
		s1.g()
	}
	if s2, ok := i.(S2); ok {
		s2.f()
		s2.g()
	}
}

func main() {
	printType(S3{})
	printType(S2{})
}
