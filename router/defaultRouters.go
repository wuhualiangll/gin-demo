package router

import (
	"gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func DefaultRouters(r *gin.Engine) {
	//根目录分组
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/login", controller.LoginController{}.Login)
		defaultRouters.GET("/news", controller.IndexController{}.Index)
		defaultRouters.GET("/index", controller.NewsController{}.GetNews)
	}

}
