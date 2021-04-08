package main

import "fmt"

func main() {
	un, in := Number(1, "第一个函数")
	fmt.Println(un, in)

	s, _ := Number(1, "第二个函数")
	fmt.Println(s)

	v, v2 := jia(2, 10)
	fmt.Println(v, v2)

	//使用不定项参数
	mo(99, "s", "v", "sss")
	ar := []string{"sdad", "dasdaq"}
	mo(100, ar...)

}

func Number(i int, v string) (a int, b int) {
	fmt.Println(v)
	return i, i
}

func jia(i int, v int) (a int, b int) {
	a = i
	b = v
	return //默认返回是这样的：return a,b
}
func mo(data int, data2 ...string) { //传入参数不确定的情况下用...声明
	for i, v := range data2 {
		fmt.Println(i, v)
	}
}
