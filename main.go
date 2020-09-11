package main

import (
	"go-demo/core"
	"go-demo/global"
	"go-demo/initialize"
)

func main() {
	//加载数据库
	switch global.GVA_CONFIG.System.DbType {
		case "mysql":
			initialize.Mysql()
		default:
			initialize.Mysql()
	}
	//注册数据库表
	initialize.DBTables()

	//程序结束前关闭数据库连接
	defer global.GVA_DB.Close()

	global.GVA_DB.LogMode(true)

	//加载redis
	initialize.Redis()

	core.RunServer()
}
