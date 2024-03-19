package main

/**
为你展示获取form表单提交的参数
并且展示参数绑定
*/
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Name     string `form:"username"` //使用标签绑定form传递过来的值
	Password string `form:"password"`
}

func main() {
	t := gin.Default()
	t.LoadHTMLFiles("./login.html")
	//页面初始话请求
	t.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	//登录请求
	t.POST("mylogin", func(c *gin.Context) {
		//	获取form传来的参数
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"你的名字": username,
			"你的密码": password,
		})
		//c.DefaultPostForm()  没获取到的话给默认值

		//使用gin框架的绑定函数,将前端传来的数据绑定到结构体
		var s UserInfo
		err := c.ShouldBind(&s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error ": "占不到数据",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"ok ": "没问题",
			})
		}
		fmt.Println(s)

	})
	t.Run(":8848")
}
