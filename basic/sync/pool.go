package main

import (
	"fmt"
	"sync"
)

type Animal struct {
	Name string
	Type string
}

func main() {
	p1 := sync.Pool{
		New: func() interface{} {
			return &Animal{}
		},
	}
	animal := p1.Get().(*Animal)
	animal.Name = "Husky"
	animal.Type = "dog"
	fmt.Println("the animal is: ", animal)
	p1.Put(animal)
	a2 := p1.Get()
	fmt.Println(a2)
}
