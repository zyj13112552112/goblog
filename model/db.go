package model

//连接数据库

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goblog/utils"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.DB,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			utils.DbUser,
			utils.DbPassword,
			utils.DbHost,
			utils.DbPort,
			utils.DbName,
		))
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}

	//禁用默认表名的复数形式
	db.SingularTable(true)

	//自动迁移模型到数据库中
	db.AutoMigrate(&User{}, &Category{}, &Article{})

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	// 这里注意不要大于gin框架的连接超时时间，因为服务器已经和客户端断开连接了
	db.DB().SetConnMaxLifetime(10 * time.Second)
}
