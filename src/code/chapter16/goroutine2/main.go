package main

import (
	"fmt"
	"runtime"
)

func main(){
	// 获取当前系统cpu的数量
	num := runtime.NumCPU()

	runtime.GOMAXPROCS(num)
	fmt.Println("num", num)

}