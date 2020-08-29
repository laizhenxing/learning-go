package main

type T4 []int

// 组合类型底层类型相同
func F4(t T4)  {

}

func main() {
	var q []int
	F4(q)
}
