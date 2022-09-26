package main

import (
	"ahmd_tools/config"
	"ahmd_tools/db"
	"ahmd_tools/web"
)

func main() {
	appConf := config.ParseConfig()

	//初始化数据库
	db.Init(appConf.Mysql)
	//启动web服务
	web.Main()
}
