package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(200, gin.H{
			"message": true,
			"name":    name,
			"action":  action,
		})
	})
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.JSON(200, gin.H{
			"firstname": firstname,
			"lastname":  lastname,
		})
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "xiaoquan")
		c.JSON(200, gin.H{
			"message": message,
			"nick":    nick,
		})
	})
	//Get + Post 混合
	//示例：
	//POST /post?id=1234&page=1 HTTP/1.1
	//Content-Type: application/x-www-form-urlencoded
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		c.JSON(200, gin.H{
			"message": message,
			"name":    name,
			"page":    page,
			"id":      id,
		})
	})

	//上传文件
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		// 上传文件到指定的路径
		c.SaveUploadedFile(file, "./test/gin/"+file.Filename)
		//返回给前段直接展示
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
		c.File("./test/gin/" + file.Filename)
	})

	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/multiUpload", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["upload"]

		for _, file := range files {
			// 上传文件到指定的路径
			c.SaveUploadedFile(file, "./test/gin/"+file.Filename)
			//返回给前段直接展示
			c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
			c.File("./test/gin/" + file.Filename)
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
