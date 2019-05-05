package main

import (
	"fmt"
	"strings"
	
)

// func AddUpper() func(int) int {
// 	var n int = 10
// 	var str string = "hello"
// 	return func(x int) int {
// 		n = n + x
// 		str += "a"
// 		fmt.Printf("str=%v", str)
// 		return n
// 	}
// }

func makeSuffix(suffix string) func (string) string{
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main(){
	// f := AddUpper()
	// fmt.Println(f(1))
	// fmt.Println(f(2))
	f := makeSuffix(".jpg")
	fmt.Println(f("hello"))
}