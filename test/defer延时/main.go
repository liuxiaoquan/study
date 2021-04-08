/*
@Time : 2021/3/8 下午2:53
@Author : 刘小全
@File : main
@Software: GoLand
*/
package main

import "fmt"

//在函数中，需要创建资源（比如：数据库连接、文件句柄、锁等等），为了在函数执行完毕后，及时的释放资源，go提供了defer机制（延时机制）
func sum(n1,n2 int) int  {
	//当执行到defer时，暂时不会执行，会将defer后面的语句ok1，ok2压入defer栈中，细节：并且会把n1和n2变量中的10，20压入栈中
	//当函数执行完毕后，再冲defer栈中，按照先入后出的方式出栈执行
	defer fmt.Println("ok1 n1=",n1)
	defer fmt.Println("ok2 n2=",n2)
	n1++
	n2++
	res := n1+n2
	fmt.Println("ok3 res=",res)
	return res
}
func main() {
	res := sum(10,20)
	fmt.Println("res=",res)
}
