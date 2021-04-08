package main

import "fmt"

func main() {
	b := func(name string, age string, v string) {
		fmt.Println(name, age)
	}
	b("刘小全", "男", "我是匿名函数")

	c := func(email string) (e string) {
		fmt.Println(email)
		return
	}
	result := c("xiaoquan@gmial.com")
	fmt.Println(result)

	(func() {
		fmt.Println("自执行函数,在程序启动中执行")
	})()

	//闭包函数调用
	re := bi("返回的值")(1)
	fmt.Println(re)

	defer diyi() //defer是延迟执行
	fmt.Println(".....")
	fmt.Println(".....")
	fmt.Println(".....")
	fmt.Println(".....")

}

//闭包函数
func bi(data string) func(data1 int) (result string) {
	return func(data1 int) (result string) {
		fmt.Println("闭包函数传入的第二个参数=", data1)
		return data
	}
}

func diyi() {
	fmt.Println("我想最先执行")
}
