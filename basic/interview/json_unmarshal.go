package main

import (
	"encoding/json"
	"fmt"
)

type P struct {
	name string `json:"name"`
}

func main() {
	js := `{
		"name": "xiao"
	}`
	//js1 := make(map[string]string)
	//js1["name"] = "test"
	var p P
	// 结构体访问控制
	// 不可访问在解析json的时候，是掠过的
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("unmarshal err: ", err)
		return
	}
	fmt.Println(p)
}
