package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimRight("CVBADBAABBBBBAAAABBAA", "BA"))	// output: ""
}
