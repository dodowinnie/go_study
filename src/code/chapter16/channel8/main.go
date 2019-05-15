package main

import (
	"fmt"
	// "time"

)



func main(){

	// 只写
	intChan := make(chan <- int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 30

	// n1 := <- intChan // 会报错
	fmt.Println(len(intChan))

	// 只读
	var redChan <-chan int
	data := <-redChan
	fmt.Println(data)


}
