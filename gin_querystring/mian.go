package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
展示 gin 获取 querystring的参数
*/
func main() {

	t := gin.Default()

	t.GET("/query", func(context *gin.Context) {
		//获取浏览器发请求携带的url参数
		//name := context.Query("query") //通过Query获取，
		name := context.DefaultQuery("query", "牛逼") //意思是查询以query为键的url中的值,查不到就给默认值
		context.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	t.Run(":8123")
}
