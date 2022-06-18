package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassWord string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败，err=", err)
		return
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8081")
}

func LoadData(file *ini.File) {

	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbName = file.Section("database").Key("DbName").MustString("blog-by-goweb")
	DbUser = file.Section("database").Key("DbUser").MustString("admin")
	DbPassWord = file.Section("database").Key("DbPassword").MustString("2222")
}
