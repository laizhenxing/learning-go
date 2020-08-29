package main

import (
	"fmt"

	"github.com/laizhenxing/laracom/demo-service/api"
)

func httpSrv() {
	fmt.Printf("Starting %v\n", appName)
	api.StartWebServer(httpPort)
}