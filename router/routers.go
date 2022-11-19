package router

import (
	"gin-demo/common"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
)

/*
*

	启动源头
*/
func SetupRouter() *gin.Engine {
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": common.UniToTime,
		"Println":    common.Println,
	})
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")
	//配置模板的文件
	r.LoadHTMLGlob("templates/**/*")
	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret11111"))
	// 设置 session 中间件，参数 mysession，指的是 session 的名字，也是 cookie 的名字
	// store 是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))
	//路由
	//admin 分组
	AdminRouters(r)
	//API 分组
	ApiRouters(r)
	//上传文件
	UploadRouters(r)
	//
	DefaultRouters(r)
	return r
}
