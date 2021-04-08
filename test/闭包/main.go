/*
@Time : 2021/2/3 下午2:33
@Author : 刘小全
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"strings"
)
//累加器
func AddUpper() func(i int) int{
	n := 10
	return func(x int) int {
		n = n + x
		return  n
	}
}
//编写一个函数makeSuffix(suffix string) 可以接收一个后缀名比如.jsg 并返回一个闭包
func makeSuffix(suffix string)func(string)string  {
	return func(fileName string) string {
		if !strings.HasSuffix(fileName,suffix) {
			return fileName + suffix
		}
		return fileName
	}
}

//用普通方式判断当前文件是否有后缀
func makeSuffix2(suffix , fileName string)string  {

	if !strings.HasSuffix(fileName,suffix) {
		return fileName + suffix
	}
	return fileName
}
func main() {
	f := AddUpper()   //返回一个闭包
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))

	fmt.Println("直接调用没有累加的效果")

	fmt.Println(AddUpper()(1))
	fmt.Println(AddUpper()(1))


	//调用makeSuffix
	ms := makeSuffix(".avi")
	fmt.Println(ms("夕阳余晖"))
	fmt.Println(ms("夕阳余晖2.avi"))

	fmt.Println(makeSuffix2(".jpg","test"))
	fmt.Println(makeSuffix2(".jpg","test2.jpg"))
}
