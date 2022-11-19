package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()

	//中间件1前设置进去，可以在中间件后的：中间件2、处理方法 获取到
	c.Set("username", "张三")

	fmt.Println("1-我是一个中间件")
	//调用该请求的剩余处理程序
	c.Next()

	fmt.Println("2-我是一个中间件")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func InitMiddlewareTwo(c *gin.Context) {

	username, _ := c.Get("username")
	fmt.Println("中间件2获取username=", username)
	fmt.Println("1-我是一个中间件-initMiddlewareTwo")
	//调用该请求的剩余处理程序
	c.Next()
	fmt.Println("2-我是一个中间件-initMiddlewareTwo")

}

func InitMiddle(c *gin.Context) {
	//定义一个goroutine统计日志  当在中间件或 handler 中启动新的 goroutine 时，不能使用原始的上下文（c *gin.Context）， 必须使用其只读副本（c.Copy()）
	cCp := c.Copy()
	//并行1
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("two! in path " + cCp.Request.URL.Path)
	}()
	//并行2
	go func() {
		fmt.Println("one " + cCp.Request.URL.Path)
	}()
}
