package controller

import (
	"gin-demo/bean"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type NewsController struct {
}

func (n NewsController) GetNews(c *gin.Context) {

	// 初始化 session 对象
	session := sessions.Default(c)
	// 通过 session.Get 读取 session 值
	username := session.Get("username")

	log.Println("session获取的值是：", username)
	news := &bean.Article{
		Title:   "新闻标题",
		Content: "新闻详情",
	}
	c.HTML(http.StatusOK, "admin/news.html", gin.H{
		"title": "新闻页面",
		"news":  news,
	})
}
