/*
@Time : 2020/12/15 下午4:30
@Author : 刘小全
@File : test
@Software: GoLand
*/
package main

import "fmt"

/**
一般程序中不主张使用goto语句
 */
func main() {
	status := 10
	fmt.Println("1")
	fmt.Println("2")
	if status==10 {
		goto la
	}
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	la :
	fmt.Println("8")
	fmt.Println("9")
	fmt.Println("10")
	fmt.Println("11")

}
