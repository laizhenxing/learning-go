package main

import "fmt"

type people struct {}

func (p *people) showA()  {
	fmt.Println("people showA")
}

func (p *people) showB()  {
	fmt.Println("people showB")
}

type teacher struct {
	people
}

func (t *teacher) showB()  {
	fmt.Println("teacher showB")
}

func main() {
	t := teacher{}
	t.showB()
}