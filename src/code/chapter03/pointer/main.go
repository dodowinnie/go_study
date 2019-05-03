package main

import(
	"fmt"
)

func main(){
	// var i int = 30
	// // i 的地址，&i
	// fmt.Println("i的地址是", &i)
	// // ptr 是一个指针变量
	// var ptr *int = &i
	// fmt.Printf("ptr=%v ptr type %T\n", ptr, ptr)
	// fmt.Println("ptr 的地址", &ptr)
	// fmt.Printf("ptr 指向的值=%v\n", *ptr)

	// var num int = 30
	// fmt.Println("num的地址是", &num)
	// var ptr *int = &num
	// fmt.Printf("ptr所对应的值是 %v", *ptr)
	// *ptr = 40
	// fmt.Printf("num的值是 %v", num)

	var n int = 10
	var ptr *int = &n
	fmt.Println("num的地址是", ptr)
	fmt.Printf("ptr所对应的值是%v", *ptr)

}