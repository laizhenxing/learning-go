package api

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Patter).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
