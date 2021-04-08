package _interface

import (
	"fmt"
)

type User struct { //实体1
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	ClassName string
	User
}

func (u User) SayName(name string) {
	fmt.Println("我的名字叫：", name)
}
