/*
@Time : 2021/3/29 上午10:22
@Author : 刘小全
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"strconv"
)
type Power struct{
	age int
	high int
	name string
}
//指针类型
func (this *Power) String() string {
	return fmt.Sprintf("age:%d, high:%d, name:%s", this.age, this.high, this.name)
}

func main() {
	strings := "abcd中国"
	fmt.Println(len(strings))

	//遍历strings中的字符
	for _,v := range strings{
		fmt.Printf("遍历完字符串之后：%c\n",v)
	}
	r := []rune(strings)
	for i:=0;i < len(r);i++ {
		fmt.Printf("fori循环：%c\n",r[i])
	}

	//字符串  转 []byte
	b := []byte("hello world!")
	fmt.Printf("byte:%v\n",b)

	//[]byte转string
	str := string(b)
	fmt.Printf("string:%v\n",str)

	//10进制转2 ， 8， 16进制
	t:=strconv.FormatInt(3421,2)
	fmt.Printf("转2进制：%v\n",t)
	h:=strconv.FormatInt(3421,8)
	fmt.Printf("转8进制：%v\n",h)
	g:=strconv.FormatInt(3421,16)
	fmt.Printf("转16进制：%v\n",g)

	var i *Power = &Power{age: 10, high: 178, name: "NewMan"}  //指针类型

	fmt.Printf("%s\n", i)
	fmt.Println(i)
	fmt.Printf("%v", i)


}
