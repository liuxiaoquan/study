package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type HelloWorld struct {
	gorm.Model
	Name string
	Sex  bool
	Age  uint
}

func main() {
	db, err := gorm.Open("mysql", "debian-sys-maint:wHjJT1VPMepXS3t3@(192.168.182.147:3306)/ginBD?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//自动检查 Product 结构是否变化，变化则进行迁移
	db.AutoMigrate(&HelloWorld{})

	//增
	//db.Create(&HelloWorld{
	//	Name:  "小米",
	//	Sex:   false,
	//	Age:   0,
	//})
	//查
	var hello HelloWorld
	db.First(&hello, "name = ?", "小红")
	fmt.Println(hello)

	var hellos []HelloWorld
	db.Where("age < ?", 25).Find(&hellos)

	fmt.Println(hellos)

	//改
	//方式一
	//db.Where("age < ?",30).Find(&HelloWorld{}).Update(&HelloWorld{
	//	Age:   18,
	//})
	//方式二
	//db.Where("name = ?","刘小全").Find(&HelloWorld{}).Update(map[string]interface{}{
	//	"Age":   18,
	//})
	//方式三
	//db.Model(&HelloWorld{}).Where("age < ?",25).Update("age",18)
	//批量修改id为1和2的属性
	//db.Where("id in (?)",[]int{1,3}).Model(&[]HelloWorld{}).Update(map[string]interface{}{
	//	"age":"26",
	//})

	//删
	db.Where("name = ?", "小红").Unscoped().Delete(&HelloWorld{})
	defer db.Close()
}
