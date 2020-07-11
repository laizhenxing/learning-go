package main

import "fmt"

func helloValue(i int)  {
	fmt.Println(i)
}

func main() {
	i := 5
	defer helloValue(i)
	i += 10

	s := []int{1,3}
	defer helloQuote(s)
	s = append(s, 4,5)
}

func helloQuote(s []int) {
	fmt.Println(s)
}
