package controller

import (
	"fmt"
	"gin-demo/bean"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (index IndexController) Index(c *gin.Context) {
	username, _ := c.Cookie("username")
	hobby, _ := c.Cookie("hobby")
	fmt.Println("获取到的Cookie的值是："+username, hobby)

	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"title": "首页",
		"msg":   " 我是msg",
		"score": 89,
		"hobby": []string{"吃饭", "睡觉", "写代码"},
		"newsList": []interface{}{
			&bean.Article{
				Title:   "新闻标题111",
				Content: "新闻详情111",
			},
			&bean.Article{
				Title:   "新闻标题222",
				Content: "新闻详情222",
			},
		},
		"news": &bean.Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		},
		"date": 1629423555,
	})
}
