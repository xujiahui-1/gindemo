package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//指定用户通过浏览器访问hello时，执行什么函数
	r.GET("/hello", sayHello)

	//启动服务
	r.Run()
}

// 被回调的函数的参数必须是 c *gin.Context
func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hellosss  ",
	})
}
