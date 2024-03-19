package main

/**
使用go解析template 示例
1.定义模板
2.解析模板
3.渲染
go get -u github.com/gin-gonic/gin
*/

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name string
	Age  int
}

// 回调的函数，必须有参数  w http.ResponseWriter, r*http.Response
func SayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("模板解析错误", err)
		return
	}
	//渲染模板 渲染入w,后面是数据
	t.Execute(w, "徐")
}

// 回调的函数，必须有参数  w http.ResponseWriter, r*http.Response
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("模板解析错误", err)
		return
	}
	//传一个结构体
	//t.Execute(w, User{Name: "xxx", Age: 13})
	// 传一个map
	myUser := User{
		Name: "徐佳辉",
		Age:  123,
	}
	myMap := map[string]interface{}{
		"user": myUser,
		"didi": "小弟弟",
	}
	t.Execute(w, myMap)
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println("Http server  start field", err)
		return
	}

}
