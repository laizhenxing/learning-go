package main

import "fmt"

type Cfg struct {
	Daemon string
}

func (c *Cfg) String() string {
	return fmt.Sprintf("print: %v", c)	// fatal error: stack overflow
}

func main() {
	c := &Cfg{}
	c.String()
}

