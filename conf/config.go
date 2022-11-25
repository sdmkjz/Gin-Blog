package conf

import (
	"fmt"
	"gin-blog/model"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode    string
	HttpPort   string
	JwtKey     string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
	BaseUrl    string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径！")
	}
	LoadServer(file)
	LoadMysql(file)
	LoadQiniu(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?&loc=Local&charset=utf8&parseTime=true"}, "")
	// 连接数据库
	model.Database(path)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("HttpPort").String()
	JwtKey = file.Section("server").Key("JwtKey").String()
	BaseUrl = file.Section("server").Key("BaseUrl").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("database").Key("Db").String()
	DbHost = file.Section("database").Key("DbHost").String()
	DbPort = file.Section("database").Key("DbPort").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").String()
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}
