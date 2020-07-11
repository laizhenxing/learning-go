package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello\n"))
	fmt.Fprintf(os.Stdout, "stdout: %#v\n", w)
	w = new(bytes.Buffer)
	fmt.Fprintf(os.Stdout, "buffer: %#v\n", w)
	w = nil
	fmt.Fprintf(os.Stdout, "nil: %#v\n", w)
}
