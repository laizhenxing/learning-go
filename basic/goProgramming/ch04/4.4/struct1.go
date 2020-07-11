package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{
		Circle: Circle{
			Point:  Point{
				X: 1,
				Y: 2,
			},
			Radius: 3,
		},
		Spokes: 4,
	}
	w.X = 4
	fmt.Printf("%#v\n", w)	// #：语法打印值，对应每个成员的名字
}
