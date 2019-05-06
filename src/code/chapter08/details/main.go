package main

import (
	"fmt"
)


func test01(arr [3]int){
	arr[0] = 88
}

func test02(arr *[3]int){
	(*arr)[0] = 88
}

func main (){
	var arr01 [3]int = [3]int{11,22,33}
	test02(&arr01)
	// var arr02 [3]float32
	// var arr03 [3]string
	// var arr04 [3]bool
	fmt.Printf("arr address:%p\n", &arr01)
	fmt.Println(arr01)
	// fmt.Println(arr02)
	// fmt.Println(arr03)
	// fmt.Println(arr04)

}