package main

import (
	"fmt"
	"time"
)

func addMsg(msg chan string) {
	msg <- "msg goroutine1"
	msg <- "msg goroutine2"
	msg <- "msg goroutine3"
	msg <- "msg goroutine4"
}

func expendMsg(msg chan string) {
	time.Sleep(2 * time.Second)
	str := <-msg
	str = str + "I'm goroutine"
	msg <- str
	close(msg)
}


func main() {
	var msg = make(chan string, 3)
	go addMsg(msg)
	go expendMsg(msg)
	time.Sleep(3 * time.Second)
	for str := range msg {
		fmt.Println(str)
	}
	fmt.Println("hello world~")
}
