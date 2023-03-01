package util

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBCharset  string
	DBLoc      string
	DBTimeout  string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
	}
	LoadServer(file)
	LoadData(file)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppModel").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadData(file *ini.File) {
	DBUsername = file.Section("database").Key("DBUsername").MustString("root")
	DBPassword = file.Section("database").Key("DBPassword").MustString("root123456")
	DBHost = file.Section("database").Key("DBHost").MustString("127.0.0.1")
	DBPort = file.Section("database").Key("DBPort").MustString("3306")
	DBName = file.Section("database").Key("DBName").MustString("diary_db")
	DBCharset = file.Section("database").Key("DBCharset").MustString("utf8mb4")
	DBLoc = file.Section("database").Key("DBLoc").MustString("Local")
	DBTimeout = file.Section("database").Key("DBTimeout").MustString("10s")

}
