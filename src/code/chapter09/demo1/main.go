package main

import (
	"fmt"
)

func main(){

	var intArr [5]int = [...]int{1,22,33,44,55}
	
	fmt.Println("arr[2] address：", &intArr[1])
	// 引用intArr数组下标包头不包尾
	slice := intArr[1:3]
	fmt.Println(slice)
	fmt.Println("slice的容量=", cap(slice)) // 切片的容量是可以动态变化的
	fmt.Println("slice address:", &slice[0])
	slice[1] = 88
	fmt.Println(intArr)
}