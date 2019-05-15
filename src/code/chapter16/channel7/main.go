package main

import (
	"fmt"
	"time"

)

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++{
		intChan <- i
	}
	close(intChan)
}

func getNum(intChan chan int, primeChan chan int, exitChan chan bool){
	var flag bool
	for{
		time.Sleep(time.Millisecond * 10)
		num, ok := <- intChan
		if !ok {
			break // 未取到
		}
		flag = true // 假设是素数
		for i := 2; i < num; i++{
			if num % i == 0 {
				// 不是素数
				flag = false
				break
			}
		}
		if flag {
			// 放进primeChan
			primeChan <- num
		}

	}
	fmt.Println("有一个primeNum协程因为取不到数据，退出")
	exitChan <- true
}

func main(){

	begin := time.Now().Unix()

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	go putNum(intChan)

	for i := 0; i < 4; i++ {
		// 开启4个协程，从intChannel取数据
		go getNum(intChan, primeChan, exitChan)
	}

	go func(){
		for i := 0; i < 4; i++ {
			// 从exitChan取出4个结果，关闭primeChan
			res := <- exitChan
			fmt.Printf("func()====%v\n",res)
		} 
		close(primeChan)
	}()

	for {
		res, ok := <- primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main线程退出")

	end := time.Now().Unix()
	fmt.Println("共计用时=", (end - begin))


}
