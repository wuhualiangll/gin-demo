package router

import (
	"gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func UploadRouters(r *gin.Engine) {
	//admin 分组
	uploadRouters := r.Group("/upload")
	{
		//先要路由到页面
		uploadRouters.GET("/add", controller.UploadController{}.Add)
		uploadRouters.GET("/addlist", controller.UploadController{}.AddLlist)

		//处理请求的方法
		uploadRouters.POST("/one", controller.UploadController{}.UploadOne)
		uploadRouters.POST("/list", controller.UploadController{}.UploadList)
	}
}
