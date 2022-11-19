package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

type UploadController struct {
}

/*
*

	先跳到这个页面
*/
func (upload UploadController) Add(c *gin.Context) {

	c.HTML(http.StatusOK, "upload/uploadone.html", gin.H{})
}

func (upload UploadController) UploadOne(c *gin.Context) {
	username := c.PostForm("username")
	//表单所绑定字段
	file, err := c.FormFile("face")

	// file.Filename 获取文件名称  aaa.jpg   ./static/upload/aaa.jpg
	dst := path.Join("./static/upload", file.Filename)
	if err == nil {
		//需要提前创建好upload 目录
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			log.Fatalln(err)
		}
	}
	// c.String(200, "执行上传")
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
}

func (upload UploadController) AddLlist(c *gin.Context) {
	c.HTML(http.StatusOK, "upload/uploadlist.html", gin.H{})
}

func (upload UploadController) UploadList(c *gin.Context) {
	username := c.PostForm("username")

	face1, err1 := c.FormFile("face1")
	dst1 := path.Join("./static/upload", face1.Filename)
	if err1 == nil {

		c.SaveUploadedFile(face1, dst1)
	}

	face2, err2 := c.FormFile("face2")
	dst2 := path.Join("./static/upload", face2.Filename)
	if err2 == nil {
		c.SaveUploadedFile(face2, dst2)
	}

	// c.String(200, "执行上传")
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst1":     dst1,
		"dst2":     dst2,
	})
}
