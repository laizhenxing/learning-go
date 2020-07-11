package main

import "fmt"

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "hello"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "hello"}
	// compare sn1 and sn2
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	// compare sm1 and sm2
	//  invalid operation: sm1 == sm2
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
