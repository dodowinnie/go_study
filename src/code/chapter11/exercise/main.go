package main

import (
	"fmt"

)

type MethodUtils struct{
	
}

func (mt MethodUtils) print(m int, n int){
	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mt MethodUtils) area(len int, width int) int{
	return len * width
} 

func main() {
	var m MethodUtils
	area := m.area(2, 4)
	fmt.Println(area)
}