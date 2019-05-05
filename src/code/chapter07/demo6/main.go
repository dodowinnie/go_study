package main

import (
	"fmt"
	"code/chapter07/demo6/utils"
)

var age = test()

func test() int{
	fmt.Println("test()....")
	return 90
}

func init(){
	fmt.Println("init().....")
}

func main(){
	fmt.Println("main().....")
	fmt.Println("age=", utils.Age, "name=", utils.Name)
}