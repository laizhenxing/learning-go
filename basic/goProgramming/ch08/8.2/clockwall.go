package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now origin: ", now)
	es, err := time.LoadLocation("US/Eastern")
	if err != nil {
		fmt.Println("time.Location error: ", err)
		os.Exit(1)
	}
	now = now.In(es)
	fmt.Println("time for US/Eastern: ", now)
}
