package web

import (
	"ahmd_tools/config"
	"ahmd_tools/web/api"
	"ahmd_tools/web/html"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Main(server config.Server) {
	if server.Port <= 0 || server.Port > 65535 {
		panic("invalid server port: " + strconv.Itoa(server.Port))
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("Hzflk@2022!"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("ahmdsession", store))

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
	r.Run(":" + strconv.Itoa(server.Port))
}
