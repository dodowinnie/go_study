package main

import (
	"fmt"
	"time"

)


func main(){
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++{
		intChan <- i
	}
	// close(intChan)
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++{
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}
	// close(strChan)

	// for v := range intChan{
	// 	fmt.Println(v)
	// }

	// for v := range strChan{
	// 	fmt.Println(v)
	// }
	// for {
	// 	v, ok := <-intChan
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }

	// for {
	// 	v, ok := <-strChan
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }

	for {
		select{
		case v:=<-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v:= <- strChan:
			fmt.Printf("从strChan读取的数据%v\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("都取不到。。。。")
			return
		}

	

	}

}
