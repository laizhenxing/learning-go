package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	_ = gin.Default()

	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err: ", err)
	}
}