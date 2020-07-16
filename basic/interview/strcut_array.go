package main

import "fmt"

type T2 struct {
	n int
}

func main() {
	ts := [2]T2{}
	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 1
			ts[1].n = 2
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
}
