package main

import (
	_ "github.com/go-sql-driver/mysql"
	"goblog/model"
	"goblog/routes"
)

/*
	项目入口
	时间：2023-1-14
	作者：郑源金
*/
func main() {
	model.InitDb()
	routes.InitRouter() //路由
}
