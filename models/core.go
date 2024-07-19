package models

import (
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var ACCOUNT, PASSWORD string
var DOMAIN, CORS string

func init() {
	cfg, _ := ini.Load("./config.ini")
	IP := cfg.Section("mysql").Key("ip").String()
	PORT := cfg.Section("mysql").Key("port").String()
	USER := cfg.Section("mysql").Key("user").String()
	PASSWORD := cfg.Section("mysql").Key("password").String()
	DATABASE := cfg.Section("mysql").Key("database").String()
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "root:root@tcp(127.0.0.1:3306)/goimage?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", USER, PASSWORD, IP, PORT, DATABASE)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Println("数据库连接失败！")
	} else {
		fmt.Println("数据库连接成功！")
	}
	AdminInit()
	DomainInit()
	EmailInit()
}

func AdminInit() {
	cfg, _ := ini.Load("./config.ini")
	ACCOUNT = cfg.Section("super").Key("account").String()
	PASSWORD = cfg.Section("super").Key("password").String()
}

func DomainInit() {
	cfg, _ := ini.Load("./config.ini")
	DOMAIN = cfg.Section("host").Key("domain").String()
	CORS = cfg.Section("host").Key("cors").String()
}

var E_ADDR, E_USERNAME, E_PASSWORD, E_HOST string

func EmailInit() {
	cfg, _ := ini.Load("./config.ini")
	E_ADDR = cfg.Section("email").Key("addr").String()
	E_USERNAME = cfg.Section("email").Key("username").String()
	E_PASSWORD = cfg.Section("email").Key("password").String()
	E_HOST = cfg.Section("email").Key("host").String()
}
