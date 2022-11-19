package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	//继承
	BaseController
}

func (u UserController) Index(c *gin.Context) {
	//源自继承
	username, _ := c.Get("username")
	fmt.Println(username)
	u.success(c)
}
func (u UserController) UserList(c *gin.Context) {
	c.String(200, "我是一个api接口-userlist")
}
func (u UserController) PList(c *gin.Context) {
	c.String(200, "我是一个api接口-plist")
}
