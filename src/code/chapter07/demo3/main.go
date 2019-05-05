package main

import (
	"fmt"
)

func main(){
	num := 20
	var ptr *int = &num
	fmt.Println("main address:", ptr)
	test(&num)
	fmt.Println("main num=", num)

}

func test(n1 *int) {
	fmt.Println("test address:", n1)
	*n1 += 10
	fmt.Println("test n1=", *n1)
}