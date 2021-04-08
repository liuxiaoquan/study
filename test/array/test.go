package main

import "fmt"

func main() {
	//跳转语句
A:
	for a := 0; a < 10; a++ {
		if a == 5 {
			break A
			goto B
		}
		fmt.Println(a)
	}
B:
	fmt.Println("我来到了B")

	//定义数组
	//方式一
	a := [3]int{1, 2, 3}
	b := [...]int{12, 3123, 324, 123, 432, 123, 435, 5654, 23}
	fmt.Println(a, b)

	//方式二
	var d = new([10]int)
	d[4] = 3
	fmt.Println(d)

	//方式三多维数组
	er := [3][3]int{ //不能写成[...][...]
		{0, 1, 2},
		{1, 2, 3},
		{2, 3, 4},
	}
	for i, v := range er {
		fmt.Println(i, v)
		for i2, v2 := range v {
			fmt.Println(i2, v2)
		}
	}

	//练习
	zoom := [...]string{"老虎", "狮子", "大象", "长颈鹿"}
	for i := 0; i < len(zoom); i++ { //第一种循环方式
		fmt.Println(zoom[i])
	}

	for i, v := range zoom { //第二种循环方式
		fmt.Println(i, v)
	}

}
