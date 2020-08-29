package main

import (
	"log"
	"net/http"

	gokithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go-kit-example/endpoints"
	"go-kit-example/transports"

	"go-kit-example/services"
)

func main() {
	srv := services.NewStringService()

	uppercaseHandler := gokithttp.NewServer(
		endpoints.MakeUppercaseEndpoint(srv),
		transports.DecodeUppercaseRequest,
		transports.EncodeResponse,
	)

	countHandler := gokithttp.NewServer(
		endpoints.MakeCountEndpoint(srv),
		transports.DecodeCountRequest,
		transports.EncodeResponse,
	)

	r := mux.NewRouter()
	r.Methods("GET").Path("/uppercase/{s}").Handler(uppercaseHandler)
	r.Methods("GET").Path("/count/{s}").Handler(countHandler)
	//http.Handle("/uppercase", uppercaseHandler)
	//http.Handle("/count", countHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
