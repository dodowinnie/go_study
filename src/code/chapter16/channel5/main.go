package main

import (
	"fmt"
	"math/rand"
	// "time"
	"strconv"
)
// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到 map 中

type Person struct{
	Name string
	Age int
	Address string
}

func test(){
	for i := 1; i <= 10; i++{
		// rand.Seed(time.Now().Unix())
		n := rand.Intn(11)
		fmt.Println(n)
	}
}

func main(){
	// test()

	var perChan chan Person
	perChan = make(chan Person, 10)

	for i := 1; i <= 10; i++ {
		// rand.Seed(time.Now().Unix())
		n := rand.Intn(11)
		// fmt.Println(n)
		person := Person{
			Name : "brandon" + strconv.Itoa(n),
			Age : n,
			Address : "shanghai" + strconv.Itoa(n),
		}
		perChan <- person
	}

	close(perChan)

	for v := range perChan {
		fmt.Printf("name=%v, age=%v,address=%v\n", v.Name, v.Age, v.Address)
	}



}
