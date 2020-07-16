package main

import "fmt"

type data struct {
	name string
}

func (d *data) print() {
	fmt.Println("name: ", d.name)
}

type printer interface {
	print()
}

func main() {
	d1 := data{
		name: "d1",
	}
	d1.print()
	fmt.Printf("%T\n", d1)

	var d2 printer = &data{"d2"}
	d2.print()
	fmt.Printf("%T\n", d2)

}
