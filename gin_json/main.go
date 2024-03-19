package main

/**
演示向前端返回json字符串
*/
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	t := gin.Default()

	//方法一：使用map
	t.GET("/jsonTest", func(c *gin.Context) {

		data := map[string]interface{}{
			"name1": "王八蛋",
			"age1":  "222",
		}
		//上面为我们自己声明的map，下面展示gin中封装的map
		//data2 := gin.H{
		//	"name": "王八蛋",
		//	"age":  "222",
		//}

		c.JSON(http.StatusOK, data)
	})

	//方法二:使用结构体
	t.GET("/jiegouti", func(context *gin.Context) {
		data := Msg{
			Name: "徐佳辉",
			Age:  13,
		}
		context.JSON(http.StatusOK, data)
	})

	t.Run(":8081")
}

type Msg struct {
	Name string
	Age  int
}
