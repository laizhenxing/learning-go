package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/bxcodec/faker/v3"
)

type Address struct {
	City string
	Area string
}

type Email struct {
	Where string `xml:"where,attr"`
	Addr string
}

type Student struct {
	Id int `xml:"id,attr"`
	Address
	Email []Email
	FirstName string `xml:"name>first"`
	LastName string `xml:"name>last"`
}

type Class struct {
	Student []Student
}

func main() {
	storeStudent()
	//deserialization()
}

func storeStudent()  {
	// 实例化对象
	stu := Student{23, Address{"shanghai","pudong"},[]Email{Email{"home","home@qq.com"}, Email{"work","work@qq.com"}},"chain","zhang"}
	store(stu)
}

func storeClass()  {
	var cla Class
	err := faker.FakeData(&cla)
	if err != nil {
		fmt.Println("fake data error: ", err)
		return
	}
	store(cla)
}

func store(data interface{})  {
	f, err := os.Create("./student.xml")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer f.Close()

	fmt.Println("student: ", data)
	// 序列化到文件中
	encoder := xml.NewEncoder(f)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("encode xml error: ", err)
		return
	}
	fmt.Println("encode xml success! Finish...")
}

