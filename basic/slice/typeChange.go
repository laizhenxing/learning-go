package main

import "fmt"

func main() {
	var t interface{} = 1.0
	if v, ok := t.(float32); ok {
		fmt.Printf("%v, %T\n", v, v)
	} else {
		fmt.Printf("change failed.\n")
	}

	//var m map[string]int
	//m["t1"] = 2
	//fmt.Println(m)

	t1 := []int{1,2}
	fmt.Println(t1[2])
}
