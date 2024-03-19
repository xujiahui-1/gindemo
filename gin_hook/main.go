package main

/**
展示z中间件钩子得使用
		1.中间件适合处理公共得业务逻辑
		2.gin.HandlerFunc
*/
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义一个中间件
func m1(ctx *gin.Context) {
	fmt.Println("m1")
	ctx.Next()  //调用后续得处理函数
	ctx.Abort() //阻止调用后续得处理函数
}
func main() {
	r := gin.Default()
	//全局注册中间件函数m1
	r.Use(m1)

	//请求处理中注册中间件函数
	r.GET("/sss", m1, func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"你好": "我不好",
		})
	})
	r.Run(":9865")
}
