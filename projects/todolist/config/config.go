package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() error {

	viper.AddConfigPath("conf")		// 设置文件路径
	viper.SetConfigName("config")	// 设置读取文件名称
	viper.SetConfigType("yaml")		// 设置文件格式
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 解析文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("解析配置文件出错：", err)
		return err
	}

	return nil
}
