package main

import "fmt"

//创建结构体
type XiaoQuan struct {
	Name   string
	Age    int
	Sex    bool
	Hobby  []string
	myHome Home
}

type Home struct {
	addr string
}

//创建结构体方法
//语法说明 fun (结构体)方法名(方法入参)(方法出参)
func (xq *XiaoQuan) Song(geName string) (restr string) {
	restr = "唱的真好"
	fmt.Printf("%v唱了一首%v,观众觉得%v\n", xq.Name, geName, restr)
	return restr
}

func (h *Home) getHome() {
	fmt.Println("我住在", h.addr)
}
func main() {
	//方式一
	var xq XiaoQuan
	xq.Name = "小全"
	xq.Age = 1
	xq.Sex = true
	xq.Hobby = []string{"aa", "bb", "cc"}

	//方式二
	xq2 := XiaoQuan{
		Name:  "小红",
		Age:   0,
		Sex:   false,
		Hobby: []string{"游泳"},
	}
	//方式三
	xq3 := XiaoQuan{
		"小米", 1, true, []string{"游泳", "跑步"}, Home{addr: "深圳"},
	}
	//方式四
	xq4 := new(XiaoQuan)
	xq4.Name = "小全"
	xq4.Age = 1
	xq4.Sex = true
	xq4.Hobby = []string{"aa", "bb", "cc"}

	//方式五，方法入参中使用struct
	quan := test(xq)
	fmt.Println(quan)

	fmt.Println(xq.Sex)
	fmt.Println(xq2.Hobby)
	fmt.Println(xq3)
	fmt.Println(xq4.Name)

	//指针，地址
	xq5 := XiaoQuan{
		Name:  "这是没有修改之前的数据",
		Age:   0,
		Sex:   false,
		Hobby: nil,
	}

	xq5x := &xq5
	xq5x.Name = "修改后"
	fmt.Println(*xq5x)

	//调用结构体方法
	qxq := XiaoQuan{
		Name:   "小红",
		Age:    0,
		Sex:    false,
		Hobby:  nil,
		myHome: Home{"深圳"},
	}
	song := qxq.Song("青藏高原")
	qxq.myHome.getHome()
	fmt.Println("\n", song)
}

func test(xq XiaoQuan) XiaoQuan {
	xq.Name = "小鹏"
	return xq
}
