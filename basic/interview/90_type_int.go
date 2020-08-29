package main

type T3 int

func F3(t T3)  {

}

func main() {
	var q int
	F3(q)	// cannot use q (type int) as type T3 in argument to F3
}
