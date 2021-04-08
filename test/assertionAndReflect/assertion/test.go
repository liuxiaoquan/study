package main

import (
	"fmt"
	_interface "yzweworkplatform-backend/test/assertionAndReflect/interface"
)

/**
断言的意思：把一个接口类型指定为它原始类型
*/
//首先我们定义两个实体，并且在实体上创建方法

func main() {
	user := _interface.User{
		Name: "刘小全",
		Age:  0,
		Sex:  false,
	}
	check(user)

	student := _interface.Student{
		ClassName: "name",
		User:      user,
	}
	check(student)
}
func check(v interface{}) {
	//v.(User).SayName(v.(User).Name)
	//断言的使用场景
	//如果该方法需要接收不确定类型的值可以用以下方法

	switch v.(type) {
	case _interface.User:
		v.(_interface.User).SayName(v.(_interface.User).Name)
		fmt.Println("我们是user")
	case _interface.Student:
		v.(_interface.Student).User.SayName(v.(_interface.Student).ClassName)
		fmt.Println("我们是student")
	}

}
