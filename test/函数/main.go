/*
@Time : 2021/1/11 下午4:21
@Author : 刘小全
@File : main
@Software: GoLand
*/
package main

import "fmt"

type mySum func(int,int) int

func sum(n1,n2 int) int  {
	return n1 + n2
}

func sum2(n1,n2 int)int  {
	return n1 - n2
}

func MyFunc(funcVar mySum,num1 int,num2 int)int  {
	return funcVar(num1,num2)
}


func main() {
	//a := sum
	//b := sum2
	fmt.Println(MyFunc(sum,1,2))
	fmt.Println(MyFunc(sum2,1,2))
}
