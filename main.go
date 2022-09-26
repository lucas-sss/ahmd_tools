package main

import (
	"ahmd_tools/db"
	"ahmd_tools/web"
)

func main() {
	//初始化数据库
	db.Init()
	//启动web服务
	web.Main()
}
