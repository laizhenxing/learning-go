package main

import "fmt"

type people1 struct {}

func (p *people1) showA()  {
	fmt.Println("people1 showA")
	p.showB()
}

func (p *people1) showB()  {
	fmt.Println("people1 showB")
}

type teacher1 struct {
	people1
}

func (t *teacher1) showB()  {
	fmt.Println("teacher1 showB")
}

func main() {
	t := teacher1{}
	t.showA()
}
