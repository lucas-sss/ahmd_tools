package web

import (
	"ahmd_tools/web/api"
	"ahmd_tools/web/html"

	"github.com/gin-gonic/gin"
)

func Main() {
	gin.SetMode(gin.ReleaseMode)
	// template目录下拥有前端用的一系列静态资源
	// box := packr.NewBox("./template")

	r := gin.Default()
	// r.StaticFS("/static", box)

	r.LoadHTMLGlob("templates/**/*")
	//指定某个目录为静态资源目录，可直接访问这个目录下的资源，url 要具体到资源名称
	r.Static("/static", "static")
	// //比前面一个多了个功能，当目录下不存 index.html 文件时，会列出该目录下的所有文件
	// r.StaticFS("/more_static", http.Dir("my_file_system"))
	// //指定某个具体的文件作为静态资源访问
	// r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	api.RegisterBusinessRoute(r)
	//注册页面路由
	html.RegistPageRoute(r)
	r.Run(":10602")
}
