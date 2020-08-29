package api

import (
	"log"
	"net/http"
)

func StartWebServer(port string)  {
	log.Println("Starting HTTP service at port: " + port)
	router := NewRouter()
	http.Handle("/", router)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Println("starting HTTP listener error: ", err.Error())
	}
}