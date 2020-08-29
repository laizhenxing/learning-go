package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var db *gorm.DB

func init()  {
	var err error
	db, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			"root", "liujing", "192.168.37.129", "3306", "gomicro"))
	if err != nil {
		log.Fatal(err)
	}

	// 最大空闲10
	db.DB().SetMaxIdleConns(10)
	// 最大连接池50
	db.DB().SetMaxOpenConns(50)
}

func GetDB() *gorm.DB {
	return db
}
