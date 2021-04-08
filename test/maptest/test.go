package main

import (
	"fmt"
	"strings"
)

func main() {
	var m map[string]string //方式一
	m = map[string]string{}

	m["name"] = "刘小全"
	m["sex"] = "男"
	fmt.Println(m)

	m1 := map[string]int{} //方式二
	m1["s"] = 1
	fmt.Println(m1)

	m2 := make(map[string]float32) //方式三
	m2["b"] = 10.2
	fmt.Println(m2)

	m3 := map[interface{}]interface{}{} //方式四
	m3["interface"] = "interface"
	m3[66] = 10.2
	m3[true] = false
	m3[4] = []int{0, 1, 2, 3}
	fmt.Println(m3)

	//map中delete操作
	delete(m3, 66)
	fmt.Println(m3)

	//map的循环
	for k, v := range m3 {
		fmt.Println("key：", k)
		fmt.Println("value：", v)
	}

	//str := "a中文cd"
	//str = string([]rune(str)[:4])
	//fmt.Println(str)

	str := "6"
	split := strings.Split(str, ",")
	for _,v := range split{
		fmt.Println(v)
	}
}
