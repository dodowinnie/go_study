package main

import (
	"fmt"
)
// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到 map 中


func main(){
	var intChan chan int
	intChan = make(chan int, 10)
	// fmt.Printf("intChan的值=%v intChan本身的地址=%p", intChan, &intChan)
	// 写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	// close(intChan)
	for i := 0; i < 10; i++ {
		<-intChan
	}
	
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	// // 从管道中读取数据
	// var num2 int
	// num2 = <- intChan
	// fmt.Println("num2=", num2)
	// fmt.Printf("channel len= %v cap=%v\n", len(intChan), cap(intChan))

	// num3 := <- intChan
	// num4 := <- intChan
	// num5 := <- intChan

	// fmt.Printf("num3=%v,num4=%v,num5=%v", num3, num4, num5)

}
