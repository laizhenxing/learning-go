package main

import "image/color"

func main() {

}

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point	// 将Point嵌入 ColoredPoint
	Color color.RGBA
}
