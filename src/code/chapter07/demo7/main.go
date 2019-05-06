package main

import (
	"fmt"
	
)

var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main(){
	// 匿名函数 方式1
	res1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(res1)

	// 匿名函数 方式2，赋值给一个变量
	a := func(n1 int, n2 int) int {
		return n2 - n1
	}

	res2 := a(10, 30)
	fmt.Println(res2)

	res3 := Fun1(10, 20)
	fmt.Println(res3)

}