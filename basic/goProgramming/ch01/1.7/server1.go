package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)	// each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

// handler echos the path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request)  {
	// 标准输出流 fmt.Fprintf
	fmt.Fprintf(w,"URL.path = %q\n", r.URL.Path)
}