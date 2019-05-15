package main

import (
	"fmt"
	"time"
	"sync"
)
// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到 map 中

var (
	myMap = make(map[int]int)
	lock sync.Mutex
)

func test(n int){
	res := 1
	for i:=1; i <= n; i++{
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main(){
	for i := 1; i <= 200; i++{
		go test(i)
	}
	time.Sleep(10*time.Second)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()

}
