package main

import (
	"fmt"
)

type myFunType func(int, int) int

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFun(funvar myFunType, num1 int, num2 int) int{
	return funvar(num1, num2)
}

func getSumAndSub(n1 int, n2 int) (sum int, sub int){
	sum = n1 + n2
	sub = n1 - n2
	return
}
func main(){
	// 函数也是一种数据类型，可以用变量接受
	a := getSum
	fmt.Printf("a的类型%T, getSum类型是%T\n", a, getSum)

	res := a(10, 40)
	fmt.Println(res)

	res2 := myFun(getSum, 10, 10)
	fmt.Println(res2)

	res3 := myFun(getSum, 20, 20)
	fmt.Println("总和：", res3)

	res4, res5 := getSumAndSub(20, 10)
	fmt.Printf("res4=%v,res5=%v", res4, res5)

}