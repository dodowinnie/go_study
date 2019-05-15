package main

import (
	"fmt"
)


func main(){

	// defer func(){
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println("程序出错,err=", err)
	// 	}
		
	// }()

	// intChan := make(chan int, 3)
	// intChan <- 100
	// intChan <- 200
	// close(intChan)

	// n1 := <- intChan
	// fmt.Println(n1)
	// intChan <- 300

	intChan := make(chan int, 100)
	for i := 0; i < 100; i++{
		intChan <- i * 2
	}
	fmt.Println(len(intChan))

	// for i := 0; i < len(intChan); i++{
	// 	data := <- intChan
	// 	fmt.Println(data)
	// }
	close(intChan)

	for v := range intChan{
		fmt.Println(v)
	}


}
