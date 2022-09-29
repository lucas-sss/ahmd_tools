package api

import (
	"ahmd_tools/config"
	"ahmd_tools/db"
	"ahmd_tools/model/apiresp"
	"ahmd_tools/uitls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserInfo struct {
	Username  string `json:"username"`
	RemoteIP  string `json:"remoteIp"`
	LoginTime string `json:"loginTime"`
}

func authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		//过滤登录页
		if uri == "/html/login" {
			c.Next()
			return
		}
		success := true
		cookie, err := c.Cookie("token")
		if err != nil {
			success = false
		} else {
			// fmt.Printf("cookie token: %s\n", cookie)
			session := sessions.Default(c)
			token := session.Get("token")
			// fmt.Printf("session token: %s\n", token)
			if token != cookie {
				success = false
			}
		}

		if success {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			//返回到登录页
			c.Redirect(http.StatusFound, "/html/login")
			// c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
	}
}

func RegisterBusinessRoute(r *gin.Engine) {

	r.POST("/business/auth/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "" {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "请输入用户名"))
			return
		}
		if password == "" {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "请输入密码"))
			return
		}
		//校验账号密码
		accounts := config.AppConf.Accounts
		if account, ok := accounts[username]; ok {
			if account.Pwd != password {
				c.JSON(http.StatusOK, apiresp.Fail(-1, "用户名或密码错误"))
				return
			}
		} else {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "用户名或密码错误"))
			return
		}
		//生成token
		uid, err := uuid.NewRandom()
		if err != nil {
			fmt.Printf("generate uuid for token fail, errmsg: %s\n", err.Error())
			c.JSON(http.StatusOK, apiresp.Fail(-1, "服务异常"))
			return
		}
		loginTime := time.Now()
		data := username + "||" + uid.String() + "||" + loginTime.String()
		fmt.Printf("token element  for account[%s], errmsg: %s\n", username, data)
		token := uitls.MD5(data)

		// 初始化session对象
		session := sessions.Default(c)
		session.Set("token", token)
		//存储登录信息
		userInfo := UserInfo{Username: username, RemoteIP: c.RemoteIP(), LoginTime: time.Now().Format("2006-01-02 15:04:05")}
		bytes, _ := json.Marshal(userInfo)
		session.Set("userInfo", string(bytes))
		session.Save()

		//设置cookie
		c.SetCookie("token", token, 1800, "/", "", false, true)
		c.JSON(http.StatusOK, apiresp.Success(nil))
	})

	r.Use(authorize())

	r.GET("/business/auth/info", func(c *gin.Context) {
		session := sessions.Default(c)
		str := session.Get("userInfo")
		userInfo := &UserInfo{}
		json.Unmarshal([]byte(str.(string)), userInfo)
		c.JSON(http.StatusOK, apiresp.Success(userInfo))
	})

	r.GET("/business/auth/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		c.SetCookie("token", "token", 0, "/", "", false, true)
		c.JSON(http.StatusOK, apiresp.Success(nil))
	})

	// /business/thridobject/search?name=弗兰科
	r.GET("/business/thridobject/search", func(c *gin.Context) {
		name := c.Query("name")

		amount, err := db.SearchObjectAmountByName(name)
		if err != nil {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "数据库查询失败"))
			return
		}
		if amount > 100 {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "查询结果条数过多，请确认企业名称，缩小搜索范围！"))
			return
		}
		objs, err := db.SearchObjectByName(name)
		if err != nil {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "数据库查询失败"))
			return
		}
		// if len(objs) > 100 {
		// 	c.JSON(http.StatusOK, apiresp.Fail(-1, "查询结果条数过多，请确认企业名称，缩小搜索范围！"))
		// 	return
		// }
		c.JSON(http.StatusOK, apiresp.Success(objs))
	})

	r.POST("/business/thridobject/pushkeyupdate", func(c *gin.Context) {
		thridobjectid := c.PostForm("thridobjectid")
		fmt.Printf("企业ID: %s\n", thridobjectid)
		if thridobjectid == "" {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "企业ID为空"))
			return
		}

		dingtalkApiResp, err := PushKeyUpdateNotice(thridobjectid)
		if err != nil {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "请求异常异常"))
			return
		}
		if dingtalkApiResp.ErrCode != 0 {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "推送密钥更新请求失败："+dingtalkApiResp.ErrorMsg))
			return
		}
		c.JSON(http.StatusOK, apiresp.Success(nil))
	})

	r.POST("/business/thridobject/boxstatus", func(c *gin.Context) {
		time.Sleep(time.Duration(2) * time.Second)
		thridobjectid := c.PostForm("thridobjectid")
		fmt.Printf("企业ID: %s\n", thridobjectid)
		if thridobjectid == "" {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "企业ID为空"))
			return
		}

		boxStatusResp, err := CheckBoxStatus(thridobjectid)
		if err != nil || boxStatusResp.Status != 0 {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "服务异常"))
			return
		}
		if boxStatusResp.Status != 0 {
			c.JSON(http.StatusOK, apiresp.Fail(-1, "查询盒子状态失败"))
			return
		}
		c.JSON(http.StatusOK, apiresp.Success(map[string]int{"state": boxStatusResp.State}))
	})
}
