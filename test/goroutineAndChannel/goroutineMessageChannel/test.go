package main

import "fmt"

/**
goroutine和channel通讯
*/
func main() {
	//定义一个channel
	c := make(chan int)
	var readc <-chan int = c
	var writec chan<- int = c
	go SetChannel(writec)
	GetChannel(readc)
}
func SetChannel(writec chan<- int) {
	for i := 0; i < 10; i++ {
		writec <- i
	}
}
func GetChannel(readc <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-readc)
	}
}
