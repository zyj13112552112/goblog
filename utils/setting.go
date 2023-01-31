package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

/*
	这个文件是读取ini配置文件，设置服务器配置参数
*/

var (
	//服务器
	AppMode  string //当前模式
	HttpPort string //服务器端口
	JwtKey   string //jwt

	//数据库
	DB         string //数据库类型
	DbHost     string //数据库主机
	DbPort     string //数据库端口
	DbUser     string //数据库用户
	DbPassword string //数据库密码
	DbName     string //数据库名字

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiNiuServer string
)

// init函数，第一次引用这个包时会调用
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取出错", err)
	}
	LoadServer(file)
	LoadDB(file)
	LoadQINIU(file)
}

func LoadServer(file *ini.File) {
	//按section->k读取v, MustString ：如果值不存在，使用这个默认值
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("202sdasdw22sada")
}

func LoadDB(file *ini.File) {
	DB = file.Section("database").Key("DB").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("42.194.222.25")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("blog")
}

func LoadQINIU(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("q_U4EeWcrJGHspQS2XwU5FPuEFqna2gza1kOcAdx")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("JQLQBs5m-1C3-sWj1pna9-2JiYK44394jO4DMXpC")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("zyjginblog")
	QiNiuServer = file.Section("qiniu").Key("QiNiuServer").MustString("http://ronyu4mwz.hn-bkt.clouddn.com/")
}
