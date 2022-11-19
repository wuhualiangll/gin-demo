package router

import (
	"gin-demo/controller"
	"gin-demo/middle"
	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	//API 分组
	apiRouters := r.Group("/api")
	{
		//调用中间间
		apiRouters.GET("/", middle.InitMiddleware, middle.InitMiddlewareTwo, controller.UserController{}.Index)
		apiRouters.GET("/userlist", middle.InitMiddle, controller.UserController{}.UserList)
		apiRouters.GET("/plist", controller.UserController{}.PList)
	}
}
