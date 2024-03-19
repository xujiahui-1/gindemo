package main

/*
	演示使用gin进行html的模板渲染
*/
import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	//创建服务
	r := gin.Default()
	//gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//静态文件的加载要写在加载模板之前
	r.Static("/xxx", "./static")
	//加载模板
	//r.LoadHTMLFiles("template/potes/index.tmpl", "template/users/index.tmpl")
	//使用正则表达式加载模板
	r.LoadHTMLGlob("template/**/*")
	//Http get请求
	r.GET("/hellopost", func(c *gin.Context) {
		//模板渲染   状态码,模板名,数据
		c.HTML(http.StatusOK, "pos/posindex", gin.H{
			//想直接向页面发送html代码并且不被转移
			"title": "<a href='http://www.baidu.com'>百度</a>",
		})
	})
	r.GET("/hellousers", func(c *gin.Context) {
		//模板渲染   状态码,模板名,数据
		c.HTML(http.StatusOK, "users/userindex", gin.H{
			"title": "徐佳辉来了",
		})
	})
	//启动服务
	r.Run(":9001")

}
