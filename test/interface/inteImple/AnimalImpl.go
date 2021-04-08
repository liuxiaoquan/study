package main

import (
	"fmt"
	"yzweworkplatform-backend/test/interface/inte"
)

var L inte.Animal //申明一个全局的接口

type Cat struct {
	Name string
	Age  string
}

func (c *Cat) Run() {
	fmt.Printf("名字叫%v的猫是%v的\n", c.Name, c.Age)
}
func (c *Cat) Eat() {
	fmt.Println(c.Name, "开始吃")
}

func MyFunc(b interface{}) {
	fmt.Println(b, "我可以接收任何值")
}

func main() {
	L = &Cat{
		Name: "Tom",
		Age:  "男",
	}
	L.Run()
	L.Eat()

	MyFunc(L)
	MyFunc([]string{"大华"})

	m := map[int]int{}
	m[1] = 1
	m[100] = 100

	for _, v := range m {
		fmt.Println(v)
	}
	MyFunc(m)

}
