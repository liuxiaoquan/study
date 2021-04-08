package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string `gorm:"primary_key;column:name;type:varchar(100);"`
}

func (u User) TableName() string {
	return "xq_user"
}
func main() {
	db, err := gorm.Open("mysql", "debian-sys-maint:wHjJT1VPMepXS3t3@(192.168.182.147:3306)/ginBD?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	defer db.Close()
}
