package main

import "fmt"

type TwoStruct1 struct {
	Name string
	Age  int
	Arr  [2]bool
	ptr  *int
	Home string
	//Email []string
	func1 func()
}

type TwoStruct2 struct {
	Name string
	Age  int
	Arr  [2]bool
	ptr  *int
	Home string
	//Email []string
	func1 func()
}

func main() {
	var s1 TwoStruct1
	var s2 TwoStruct2
	s3 := TwoStruct2(s1)
	fmt.Println("s3 == s2", s3 == s2)
}
