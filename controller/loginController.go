package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
}

func (l LoginController) Login(c *gin.Context) {
	username, _ := c.GetQuery("username")
	//过期时间延时
	c.SetCookie("hobby", "吃饭 睡觉", 5, "/", "localhost", false, true)
	fmt.Println(username, "=====username")
	c.SetCookie("username", username, 3600, "/", "localhost", false, true)

	//sessions
	//初始化 session 对象
	session := sessions.Default(c)
	//设置过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, // 6hrs
	})
	//设置 Session
	session.Set("username", username)
	session.Save()
	c.String(http.StatusOK, username)
}
