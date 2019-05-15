package main

import (
	"fmt"
	"time"

)

func WriteData(intChan chan int){
	// 写数据
	for i := 1; i <= 10; i++ {
		intChan <- i
		fmt.Println("writeData ", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func ReadData(intChan chan int, exitChan chan bool){
	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

func main(){

	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go WriteData(intChan)
	go ReadData(intChan, exitChan)

	for {
		v, ok := <- exitChan
		if !ok {
			fmt.Printf("v=%v,ok=%v", v, ok)
			break
		}
	}


}
