package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
	//"yzweworkplatform-backend/model"
)

type Class struct {
	gorm.Model
	CLassName string
	Student   []Student
}
type Student struct {
	gorm.Model
	Name    string
	ClassID uint
	IDCard  IDCard
	//多对多
	Teacher []Teacher `gorm:"many2many:student_teacher"`
}
type IDCard struct {
	gorm.Model
	StudentID uint
	ADDR      string
}
type Teacher struct {
	gorm.Model
	//多对多
	TeacherName string
	Student     []Student `gorm:"many2many:student_teacher"`
}

func main() {
	//db, err := gorm.Open("mysql", "debian-sys-maint:wHjJT1VPMepXS3t3@(192.168.182.147:3306)/ginBD?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "debian-sys-maint:wHjJT1VPMepXS3t3@(192.168.182.147:3306)/qmPlus?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//自动检查 Product 结构是否变化，变化则进行迁移
	//db.AutoMigrate(
	//	&Class{},
	//	&Student{},
	//	&IDCard{},
	//	&Teacher{},
	//	)
	//defer db.Close()

	//i := IDCard{
	//	//StudentID: 0,
	//	ADDR:      "深圳市",
	//}
	//s := Student{
	//	Name:    "刘小全",
	//	IDCard:  i,
	//	//ClassID: 1,
	//	Teacher: nil,
	//}
	//
	//t := Teacher{
	//	TeacherName: "老师傅",
	//	Student:   []Student{s},
	//}
	//c := Class{
	//	CLassName: "一班",
	//	Student:   []Student{s},
	//}
	//_=db.Create(&c).Error
	//_=db.Create(&t).Error
	r := gin.Default()
	r.POST("/student", func(c *gin.Context) {
		var student Student
		c.BindJSON(&student)

		var res Student
		//判断是否存在该学生
		db.Where("name = ?", student.Name).First(&res)
		if res.Name == "" {
			db.Create(&student)
			c.JSON(200, gin.H{
				"message": "插入成功",
				"student": student,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "已经注册,插入失败",
			})
		}

	})
	r.GET("/student/:id", func(c *gin.Context) {
		id := c.Param("id")
		var student Student
		db.Preload("Teacher").Preload("IDCard").
			Where("id = ?", id).FirstOrInit(&student, Student{
			Name: "没有查到数据",
		})
		c.JSON(200, gin.H{
			"student": student,
		})
	})
	//分页查询
	r.GET("/teacher/:page/:index", func(c *gin.Context) {
		page := c.Param("page")
		index := c.Param("index")
		var teacher []Teacher
		pageunm, _ := strconv.Atoi(page)
		indexunm, _ := strconv.Atoi(index)
		//分页数据
		db.Limit(indexunm).Offset(indexunm * (pageunm - 1)).Find(&teacher)
		//总数
		var count int
		db.Find(&teacher).Count(&count)
		if len(teacher) == 0 {
			c.JSON(202, gin.H{
				"message": "最后一页了",
			})
		} else {
			c.JSON(200, gin.H{
				"count":   count,
				"page":    page,
				"index":   index,
				"teacher": teacher,
			})
		}
	})
	r.GET("/class/:id", func(c *gin.Context) {
		id := c.Param("id")
		var class Class
		db.Preload("Student").Preload("Student.IDCard").Preload("Student.Teacher").Find(&class, "id = ?", id)
		c.JSON(200, gin.H{
			"class": class,
		})
	})

	r.PUT("/update/:id", func(c *gin.Context) {
		id := c.Param("id")
		var class Class
		db.Model(&class).Where("id = ?", id).Update("CLassName", "二班")
		c.JSON(200, gin.H{
			"class": class,
		})
	})

	//r.GET("/wxUser/:id", func(c *gin.Context) {
	//	id := c.Param("id")
	//	var wxUser model.WeworkUser
	//	db.Find(&wxUser, "id = ?", id)
	//
	//	var wxUserDepartment []model.WeworkUserDepartment
	//	db.Where("user_id = ?", wxUser.ID).Find(&wxUserDepartment)
	//
	//	var data []uint
	//	for _, v := range wxUserDepartment {
	//		data = append(data, v.DepartmentIds)
	//	}
	//	var wxDepartment []model.WeworkDepartment
	//	db.Where("id in (?)", data).Find(&wxDepartment)
	//	wxUser.Department = wxDepartment
	//
	//	c.JSON(200, gin.H{
	//		"wxUser": wxUser,
	//	})
	//})
	r.Run(":8888")

}
