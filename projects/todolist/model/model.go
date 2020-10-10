package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/spf13/viper"
)

type Model struct {
	ID uint `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT;" json:"id"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"deleted_at" json:"deleted_at"`
}

var DB *gorm.DB

func InitDB() error {
	var err error
	conf := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.db_name"))
	DB, err = gorm.Open("mysql", conf)
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return err
	}

	return nil
}