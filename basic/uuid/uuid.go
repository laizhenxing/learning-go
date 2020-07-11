package main

import (
	"fmt"

	uid "github.com/satori/go.uuid"
	guid "github.com/google/uuid"
)

func main() {
	// 创建
	u1 := uid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

	// 解析
	u2, err := uid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Printf("something wrong: %s\n", err)
		return
	}
	fmt.Printf("successfully~ %s\n", u2)

	gu1 := guid.New()
	fmt.Printf("guid：%s\n", gu1)

	gu2, err := guid.Parse("7360d4e7-9f74-4c24-abcc-fe200fe4c771")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("gu2: %s", gu2)
}
