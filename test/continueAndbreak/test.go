package main

import "fmt"

func main() {
	for i:=0;i<=10;i++ {
		fmt.Println("i=",i)
		break
	}

	for k:=0;k<=10;k++ {
		fmt.Println("k=",k)
		continue
		fmt.Println("continue后的代码")
	}

	here:
		for i:=0;i<2;i++ {
			for j:=1;j<4;j++ {
				if j==2 {
					continue here
				}
				fmt.Println("i=",i,"j=",j)
			}
		}
}
