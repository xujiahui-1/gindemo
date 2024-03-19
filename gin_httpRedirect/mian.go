package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
向您展示http 的内部重定向,外部重定向
*/
func main() {

	t := gin.Default()

	t.GET("/redirect", func(context *gin.Context) {
		//	重定向
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	t.GET("/a", func(context *gin.Context) {
		//	跳转到b控制器 ,请求转发
		context.Request.URL.Path = "/b" //把请求的地址进行修改
		t.HandleContext(context)
	})
	t.GET("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"你好": "我不好",
		})
	})
	t.Run(":8745")
}
