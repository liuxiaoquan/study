package main

import (
	"fmt"
)

func main() {
	var level int = 9
	for  i := 1; i<= level ;i++ {
		for k :=1; k <= level - i; k++ {
			fmt.Print(" ")
		}
		for j := 1;j <= 2 * i -1;j++ {
			if j==1 || j == 2 * i -1 || i == level{
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}


	triangle()
	for i := 1;i <= level; i++ {
		for j := 1;j<=i;j++ {
			fmt.Printf("%v * %v = %v \t",j,i,j*i)
		}
		fmt.Println()
	}
}

func triangle()  {
	level := 9
	for i :=1 ;i<=2 * level -1; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	for i :=1; i<level;i++ {
		for j:= 0;j<2 * level - 1 - i;j++ {
			if j==2 * level - i - 2 || j==i {
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
