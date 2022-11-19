package router

import (
	"gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	//admin 分组
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/index", controller.IndexController{}.Index)
		adminRouters.GET("/news", controller.NewsController{}.GetNews)
	}
}
