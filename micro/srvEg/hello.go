package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello micro"))
	})

	http.ListenAndServe(":9000", nil)
}
