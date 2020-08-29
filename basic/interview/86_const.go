package main

import "fmt"

const (
	Century = 100
	// 以 0 开头，8进制
	Decade = 010	// 1*8^1+0*8^0
	Year = 001	// 1*8^0
)

func main() {
	fmt.Println(Century)
	fmt.Println(Decade, 2*Decade)
	fmt.Println(Year, 2*Year)
	fmt.Println(Century + 2*Decade + 2*Year)
}
