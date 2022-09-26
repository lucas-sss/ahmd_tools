package html

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistPageRoute(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		// 指定重定向的URL
		c.Redirect(http.StatusFound, "/html/login")
	})

	r.GET("/html/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login/login.html", nil)
	})

	r.GET("/html/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.html", nil)
	})
}
