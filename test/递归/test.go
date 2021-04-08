/*
@Time : 2020/12/17 上午10:56
@Author : 刘小全
@File : test
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

func test(data int)  {
	if data > 2 {
		data--
		test(data)
	}
	fmt.Println("test() data=",data)

}

func fblq (n int) int{
	if n==1 || n==2 {
		return 1
	}else {
		 data1 := fblq(n - 1)
		 data2 := fblq(n - 2)
		 result := data1+data2
		return result
	}
}
func hanshu(n int)int  {
	if n == 1 {
		return 3
	}else {
		return 2 * hanshu(n - 1) + 1
	}
}

func getPartID(node []uint) (partIds []uint) {
	for _, v := range node { //一层
		partIds = append(partIds, v)
		if v != 0 { //判断是否为空
			node = node[1:]
			getPartID(node)
		}
	}
	return partIds
}

/**
 * @Author lxq
 * @Description //TODO 有一堆桃子，猴子第一天吃了其中的一半，并且再多吃了一个！以后每天猴子都吃其中的一半，然后再多吃一个，当第10天时，想再吃时（还没吃），发现只有一个桃子了。问题：最初一共有多少个桃子？
 * @Date 下午3:24 2021/1/5
 * @Param
 * @return
 **/
func eat(n int)int  {
	if n == 10 {
		return 1
	}else {
		return (eat(n + 1) + 1)  * 2
	}
}

func test01(n1 *int){
	fmt.Println(n1)
	*n1 = *n1 + 10
	fmt.Println(n1)
}
func main() {
	var node []uint
	node = append(node, 5)
	node = append(node, 10)
	node = append(node, 10)
	node = append(node, 10)
	node = append(node, 10)
	fmt.Println(getPartID(node))

	fmt.Println("第3个斐波那契数列的值为",fblq(3))
	fmt.Println("第4个斐波那契数列的值为",fblq(4))
	fmt.Println("第5个斐波那契数列的值为",fblq(5))
	fmt.Println("第6个斐波那契数列的值为",fblq(6))
	fmt.Println("第7个斐波那契数列的值为",fblq(7))

	fmt.Println("计算出来的函数为：",hanshu(5))
	fmt.Println("10天猴子的桃子数量为：",eat(8))
	num := 10
	test01(&num)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}


