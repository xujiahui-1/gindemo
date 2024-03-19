package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
展示gin中的路由于路由组
*/
func main() {
	t := gin.Default()
	//t.GET t.POST t.PUT  t.DELETE
	t.GET("/route", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"你好": "我不好",
			"我好": "你不好",
		})
	})
	//NoRoute用来处理404基本
	t.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "404",
		})
	})

	//定义路由组,把统一的路径统合起来
	//group := t.Group("/user")
	//group.GET()
	//group.GET()
	//group.GET()
	//...
	t.Run(":8858")
}
