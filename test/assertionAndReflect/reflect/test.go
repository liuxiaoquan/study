package main

import (
	"fmt"
	"reflect"
)

//反射通俗来讲，通过反射就能操作原始的数据类型，方法，或者说可以操作它原始的值
func main() {
	//u := _interface.User{
	//	Name: "小全",
	//	Age:  0,
	//	Sex:  false,
	//}
	//check(u)

	v := []string{"ss"}
	check(v)
}

func check(inter interface{}) {
	t := reflect.TypeOf(inter)         //获取值的类型
	v := reflect.ValueOf(inter)        //获取参数种数据的值
	of := reflect.TypeOf(inter).Kind() //用来判断类型

	fmt.Println(t, v, of)
}
