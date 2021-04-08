package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func milled() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("方法前")
		c.Next()
		fmt.Println("方法后")
	}
}

func logs(param gin.LogFormatterParams) string {
	// 你的自定义格式
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func main() {

	// 创建记录日志的文件
	f, _ := os.Create("./test/gin/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	v1 := r.Group("/v1").Use(milled()).Use(gin.LoggerWithFormatter(logs))
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
			fmt.Println("---------进入方法1------------")
		})
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
			fmt.Println("---------进入方法2------------")
		})
	}
	r.Run(":8081")
}
