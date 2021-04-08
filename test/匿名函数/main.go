/*
@Time : 2021/1/28 下午4:22
@Author : 刘小全
@File : main
@Software: GoLand
*/
package main

import "fmt"

var(
	Funcs = func (n,n1 int)int { //全局匿名函数
		return n1*n
	}
)
func main() {
	//匿名函数只调用一次
	result :=func (n1,n2 int) int{
		return n1+n2
	}(1,8)
	fmt.Println(result)

	//匿名函数多次调用
	funcResult :=func (n1,n2 int) int{
		return n1 - n2
	}
	result = funcResult(1,2)
	fmt.Println(result)
	fmt.Println("全局匿名函数：",Funcs(3,2))
}
