package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 初始化一个http服务对象
	app := gin.Default()
	// 首先加载templates目录下面的所有模版文件，模版文件扩展名随意
	app.LoadHTMLGlob("templates/*")
	app.Static("/static", "./static")
	// 设置一个get请求的路由，url为/ping, 处理函数（或者叫控制器函数）是一个闭包函数。
	app.GET("/ping", func(context *gin.Context) {
		// 通过请求上下文对象Context, 直接往客户端返回一个json
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	app.GET("/", func(context *gin.Context) {
		// 通过HTML函数返回html代码
		// 第二个参数是模版文件名字
		// 第三个参数是map类型，代表模版参数
		// gin.H 是map[string]interface{}类型的别名
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	app.Run(":45678") // 监听并在 0.0.0.0:8080 上启动服务

}
