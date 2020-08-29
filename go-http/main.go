package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "<h1>hello world</h1>\n")
}

func main() {
	//search()
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}

func search()  {
	a := []int{55, 45, 36, 28, 21, 15, 10, 6, 3, 1}
	x := 6
	// 使用二分法
	i := sort.Search(len(a), func(i int) bool { return a[i] <= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}

