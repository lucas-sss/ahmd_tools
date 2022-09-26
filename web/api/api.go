package api

import (
	"ahmd_tools/db"
	"ahmd_tools/model/response"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterBusinessRoute(r *gin.Engine) {

	r.POST("/business/auth/login", func(c *gin.Context) {
		//TODO
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// /business/thridobject/search?name=弗兰科
	r.GET("/business/thridobject/search", func(c *gin.Context) {
		name := c.Query("name")
		objs, err := db.SearchObjectByName(name)
		if err != nil {
			c.JSON(http.StatusOK, response.Fail(-1, "数据库查询失败"))
			return
		}
		if len(objs) > 100 {
			c.JSON(http.StatusOK, response.Fail(-1, "查询结果条数过多，请精确查询条件"))
			return
		}
		c.JSON(http.StatusOK, response.Success(objs))
	})

	r.POST("/business/thridobject/pushkeyupdate", func(c *gin.Context) {
		time.Sleep(time.Duration(2) * time.Second)
		thridobjectid := c.PostForm("thridobjectid")
		fmt.Printf("企业ID: %s\n", thridobjectid)
		if thridobjectid == "" {
			c.JSON(http.StatusOK, response.Fail(-1, "企业ID为空"))
			return
		}
		c.JSON(http.StatusOK, response.Success(nil))
	})

	r.POST("/business/thridobject/boxstatus", func(c *gin.Context) {
		time.Sleep(time.Duration(2) * time.Second)
		thridobjectid := c.PostForm("thridobjectid")
		fmt.Printf("企业ID: %s\n", thridobjectid)
		if thridobjectid == "" {
			c.JSON(http.StatusOK, response.Fail(-1, "企业ID为空"))
			return
		}
		c.JSON(http.StatusOK, response.Success(nil))
	})
}
