package main

import (
	"fmt"
)
// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到 map 中

type Cat struct{
	Name string
	Age int
}

func main(){
	// var intch chan int
	// intch = make(chan int, 3)
	// intch <- 10
	// intch <- 20
	// intch <- 10

	// n1 := <- intch
	// n2 := <- intch
	// n3 := <- intch

	// var mapch chan map[string]string
	// mapch = make(chan map[string]string, 10)
	// m1 := make(map[string]string)
	// m1["city1"] = "北京"
	// m1["city2"] = "天津"

	// m2 := map[string]string{
	// 	"hero1":"宋江",
	// 	"hero2":"武松",
	// }

	// mapch <- m1
	// mapch <- m2

	// mm1 := <- mapch
	// mm2 := <- mapch

	// fmt.Printf("num1=%v,num2=%v, num3=%v", n1, n2, n3)
	// fmt.Printf("mm1=%v,mm2=%v", mm1, mm2)

	// var catChan chan Cat
	// catChan = make(chan Cat, 10)

	cat1 := Cat{
		Name:"tom",
		Age:10,
	}

	// cat2 := Cat{
	// 	Name:"brandon",
	// 	Age:20,
	// }
	// catChan <- cat1
	// catChan <- cat2

	// cat11 := <- catChan
	// cat22 := <- catChan
	// fmt.Println(cat11, cat22)

	// var catptrChan chan *Cat
	// catptrChan = make(chan *Cat, 10)

	// cat1 := &Cat{
	// 	Name:"tom",
	// 	Age:10,
	// }

	// cat2 := &Cat{
	// 	Name:"brandon",
	// 	Age:20,
	// }

	// fmt.Printf("cat1=%p,cat2=%p\n", cat1, cat2)

	// catptrChan <- cat1
	// catptrChan <- cat2

	// cat11 := <- catptrChan
	// cat22 := <- catptrChan
	// fmt.Println(cat11, cat22)

	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	allChan <- 10
	allChan <- 'a'
	allChan <- "brandon"
	allChan <- 10.25
	allChan <- true
	allChan <- cat1

	data1 := <- allChan
	data2 := <- allChan
	data3 := <- allChan
	data4 := <- allChan
	data5 := <- allChan
	data6 := <- allChan
	cat := data6.(Cat) // 使用类型断言

	fmt.Printf("data1=%v\n", data1)
	fmt.Printf("data2=%v\n", data2)
	fmt.Printf("data3=%v\n", data3)
	fmt.Printf("data4=%v\n", data4)
	fmt.Printf("data5=%v\n", data5)
	fmt.Printf("data6=%v\n", cat.Name)




}
