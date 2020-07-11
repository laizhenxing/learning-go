package main

import "fmt"

func main() {
	s1()
	s2()
}

func s1()  {
	s1 := make([]int, 5)
	s1 = append(s1,1,2,3,4)
	fmt.Println("func s1: ", s1)
}

func s2()  {
	s := make([]int, 0)
	s = append(s,1,2,3,4)
	fmt.Println("func s2:", s)
}

func funcMui(x,y int) (sum int, err error) {
	return x+y, nil
}
