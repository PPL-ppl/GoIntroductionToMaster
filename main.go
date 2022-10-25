package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"log"
	"strconv"
)

// 自定义Go中间件-拦截器

func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("userSession", "userid-1")
		context.Next() //放行
		//context.Abort() //拦截
	}
}

func main() {
	//创建一个服务
	ginServer := gin.Default()
	//注册自定义中间件
	ginServer.Use(myHandler())
	ginServer.Use(favicon.New("./favicon.ico"))
	//访问地址，处理Request请求 Response响应
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello world"})
	})
	ginServer.GET("/helloStudent", func(context *gin.Context) {
		context.JSON(200, Student{Name: "小刘", Id: 22})
	})

	//加载静态页面
	ginServer.LoadHTMLGlob("templates/*")
	//加载资源文件
	ginServer.Static("static", "./static")
	//响应静态页面
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", Student{Name: "小刘", Id: 22})
	})
	//接收前端传过来的参数
	//http://127.0.0.1:8088/user/info?Id=2&Name=2
	ginServer.GET("/user/info", func(context *gin.Context) {
		id := context.Query("Id")
		name := context.Query("Name")
		i, _ := strconv.Atoi(id)
		student := Student{i, name}
		context.JSON(200, student)
	})
	//http://127.0.0.1:8088/user/info/1/2
	ginServer.GET("/user/info/:Id/:Name", func(context *gin.Context) {
		id := context.Param("Id")
		name := context.Param("Name")
		i, _ := strconv.Atoi(id)
		student := Student{i, name}
		context.JSON(200, student)
	})

	//前端给后端传Json
	ginServer.POST("/json", func(context *gin.Context) {
		data, _ := context.GetRawData()
		var student Student
		_ = json.Unmarshal(data, &student)
		context.JSON(200, student)
	})
	//表单提交  支持函数式编程
	ginServer.POST("/json/info", func(context *gin.Context) {
		id := context.PostForm("Id")
		name := context.PostForm("Name")
		context.JSON(200, gin.H{
			"msg":  "ok",
			"Id":   id,
			"name": name,
		})
	})
	//路由
	ginServer.GET("/test", func(context *gin.Context) {
		//重定向301
		context.Redirect(301, "https://www.baidu.com")
	})
	// 404 NoRoute
	ginServer.NoRoute(func(context *gin.Context) {
		//重定向404
		context.HTML(404, "index.html", nil)
	})
	//路由组
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.GET("/delete")
		userGroup.GET("/logout")
	}
	//使用中间件  拦截器  前面实现了自定义中间件

	ginServer.GET("/getUser", myHandler(), func(context *gin.Context) {
		s := context.MustGet("userSession").(string)
		log.Println(s)
	})

	//服务器端口
	ginServer.Run(":8088")
}

type Student struct {
	Id   int
	Name string
}
