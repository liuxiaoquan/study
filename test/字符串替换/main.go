/*
@Time : 2021/2/6 下午9:41
@Author : 刘小全
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str :=`{"Atype":"hash", "ts":1612616191727, "data":[{"Atype":"hash", "vtype":"alertbox", "title":"温馨提示", "content":"hohohoho", "alerttype":"info", "okcmd":"", "oktitle":"", "cancelcmd":"", "canceltitle":"", "zKeyOrder":["vtype", "title", "content", "alerttype", "okcmd", "oktitle", "cancelcmd", "canceltitle"]}], "userid":"", "action":"chat", "code":0, "msg":"", "zKeyOrder":["ts", "data", "userid", "action", "code", "msg"]}`

	fmt.Println(ReplaceStr(str,``))
}

func ReplaceStr(str,newStr string) (replaceStr string ) {
	reg := regexp.MustCompile(`"A([^"]*)":"h([^"]*)",\s`)
	results := reg.FindAllString(str, -1)
	for _,v := range results{
		replaceStr = strings.Replace(str, v, newStr, -1)
	}
	reg = regexp.MustCompile(`,\s"zKe([^\]]*\])`)
	results = reg.FindAllString(replaceStr, -1)
	for _,v :=range results{
		replaceStr = strings.Replace(replaceStr, v, newStr, -1)
	}
	return replaceStr
}