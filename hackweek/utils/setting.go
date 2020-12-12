package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	JwtKey   string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)

}

func LoadServer(file *ini.File) {
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}


