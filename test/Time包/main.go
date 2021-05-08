/*
@Time : 2021/3/30 下午2:36
@Author : 刘小全
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	today_str := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("%v\n", today_str)

	st := fmt.Sprint(now.Format("2006/01/02 15:04:05"))
	fmt.Println(st)
	st = fmt.Sprint(now.Format("2006-01-02 15:04:05"))
	fmt.Println(st)
	st = fmt.Sprint(now.Format("2006-01-02"))
	fmt.Println(st)
	st = fmt.Sprint(now.Format("2006"))
	fmt.Println(st)
	st = fmt.Sprint(now.Format("01"))
	fmt.Println(st)
	st = fmt.Sprint(now.Format("1"))
	fmt.Println(st)

	startTimeM := time.Now().Unix()
	startTimeN := time.Now().UnixNano()
	test()
	endTimeM := time.Now().Unix()
	endTimeN := time.Now().UnixNano()
	fmt.Printf("执行test花费毫秒数为：%v\n", endTimeM-startTimeM)
	fmt.Printf("执行test花费纳秒数为：%v\n", endTimeN-startTimeN)
}

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
