package main

import "fmt"

func main() {
	//channal是goroutine之间的通讯桥梁
	//定义channal分五种

	//第一种：有缓冲
	c1 := make(chan int, 5) //定义一个chan，有缓冲，有五个缓冲
	c1 <- 1                 //在缓冲区存一个int类型的值
	fmt.Println(<-c1)       //在缓冲区取一个int类型的值

	//第二种：无缓冲区
	c2 := make(chan int) //无缓冲
	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i //循环在chan里面保存值
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c2) //循环取值
	}
	//第五种：可读可取
	c3 := make(chan int, 5)
	//第三种：可读缓冲
	var read <-chan int = c3
	//第四种：可写缓冲
	var write chan<- int = c3
	write <- 199
	fmt.Println(<-read)

	//chan的close用法
	var cc = make(chan int, 10)
	cc <- 1
	cc <- 2
	cc <- 3
	cc <- 4
	cc <- 5
	close(cc) //如果没有关闭会报异常
	for i := 0; i < 5; i++ {
		fmt.Println(<-cc)
	}

	//chan的select用法
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 1)
	chan3 := make(chan int, 1)
	chan1 <- 1
	chan2 <- 2
	chan3 <- 3
	select {
	case <-chan1:
		fmt.Println("chan1")
	case <-chan2:
		fmt.Println("chan2")
	case <-chan3:
		fmt.Println("chan3")
	}

}
