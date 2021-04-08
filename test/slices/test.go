package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2}
	cl := a[:] //切片
	cl2 := a[1:]
	cl3 := a[:2]
	fmt.Println("cl=", cl, "cl2=", cl2, "cl3=", cl3)
	cl = append(cl, 10)
	fmt.Println("append之后的cl=", cl)
	copy(cl, cl2)
	fmt.Println("copy之后的cl=", cl)
	copy(cl[1:], cl2)
	fmt.Println("copy标注下标之后的cl=", cl)

	//切片就是一个没有固定长度的数值类似与Java中的arrayList
	//如何声明一个切片
	var qiepian []int //第一个声明方式
	qiepian = append(qiepian, 1, 2, 3, 4, 5)
	fmt.Println("声明切片的效果=", qiepian)

	for i, v := range qiepian {
		fmt.Println("下标为：", i)
		fmt.Println("值为：", v)
	}
	for i := 0; i < len(qiepian); i++ {
		fmt.Println(qiepian[i])
	}
	//切片第二个声明方式
	qia := make([]float32, 6)
	fmt.Println(qia)
}
