package main

import (
	"fmt"

	"gokit/util"
)

func main() {
	err := util.GenRsaKey(1024, "pems")
	if err != nil {
		fmt.Println(err)
	}
}
