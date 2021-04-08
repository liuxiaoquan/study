package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name" binding:"required"`
	Age uint32 `json:"age"`
	Addr string `json:"addr"`
}
func main() {
	r := gin.Default()
	r.POST("/user/getName", func(c *gin.Context) {
		var p Person
		c.ShouldBindJSON(&p)
		c.JSON(200, gin.H{
			"message": true,
			"name":    p.Name,
			"age":  p.Age,
			"addr": p.Addr,
		})
	})
	port := flag.String("port", ":8080", "usage http web server port number")
	flag.Parse()
	r.Run(*port) // listen and serve on 0.0.0.0:8080
	//go build test.go打包之后可以用./test -port=:9999 来指定端口
}
