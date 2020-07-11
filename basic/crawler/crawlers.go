package main

import (
	"crawler/helpers"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reEmail = `\w+@\w+\.\w+`
	rePhone = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIdCard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImg = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
	// +?代表贪婪模式
	reLink = `href="(https?://[\s\S]+?)"`
)

func main()  {
	//email.GetEmail()
	//users.GetUsers()
	// 获取email
	//GetEmail("https://studygolang.com/users")
	// 获取链接
	//GetLinks("https://studygolang.com/")
	// 获取身份证
	//GetIdCard("https://henan.qq.com/a/20171107/069413.htm")
	// 获取图片连接
	GetImages("https://blog.csdn.net/qq_20432379/article/details/86506432")
	// 获取手机号
	//GetPhone("https://www.zhaohaowang.com/")
}

// 获取url的内容
func GetPageStr(url string) (pageStr string) {
	rsp, err := http.Get(url)
	helpers.HandlerError(err, "http.Get url")
	defer rsp.Body.Close()
	// 获取网页内容
	pageBytes, err := ioutil.ReadAll(rsp.Body)
	helpers.HandlerError(err, "ioutil.ReadAll")
	pageStr = string(pageBytes)
	return
}


func GetEmail(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reEmail)
	res := re.FindAllStringSubmatch(pageStr, -1)
	for _, r := range res {
		fmt.Println(r)
	}
}

// 获取连接
func GetLinks(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLink)
	res := re.FindAllStringSubmatch(pageStr, -1)
	for _, r := range res {
		fmt.Println(r)
	}
}

// 获取身份证号
func GetIdCard(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reIdCard)
	res := re.FindAllStringSubmatch(pageStr, -1)
	for _, r := range res {
		fmt.Println(r)
	}
}

// 获取电话号码
func GetPhone(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhone)
	res := re.FindAllStringSubmatch(pageStr, -1)
	for _, r := range res {
		fmt.Println(r)
	}
}

// 获取图片
func GetImages(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	res := re.FindAllStringSubmatch(pageStr, -1)
	for _, r := range res {
		fmt.Println(r)
	}
}
