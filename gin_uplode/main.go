package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
展示gin实现文件上传
*/
func main() {
	t := gin.Default()
	t.LoadHTMLFiles("./index.html")
	t.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	t.POST("/myupload", func(context *gin.Context) {
		//	接受文件
		file, err := context.FormFile("f1")
		if err != nil {
			context.JSON(400, gin.H{
				"error": "文件上传失败",
			})
		} else {
			//将读取到的文件保存在本地
			dst := fmt.Sprintf("./%s", file.Filename)
			context.SaveUploadedFile(file, dst)
		}

	})

	t.Run(":8787")
}
