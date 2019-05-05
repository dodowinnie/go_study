package main

import(
	"fmt"
	util "code/chapter07/demo1/utils"
)

func main() {

	fmt.Println("utils包中的变量：", util.Num)
	util.SayHello()
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '-'
	res := util.Cal(n1, n2, operator)
	fmt.Println(res)
}

