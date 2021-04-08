package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup //携程管理器
	wg.Add(1)
	go Run(&wg)
	//time.Sleep(1*time.Second)
	wg.Wait()

}

func Run(wg *sync.WaitGroup) {
	fmt.Println("我跑起来了")
	wg.Done()
}
