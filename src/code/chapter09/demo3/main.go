package main

import (
	"fmt"
)

func main(){
	str := "hello@brandon"
	slice := []rune(str)
	slice[0] = '北'
	str = string(slice)
	// arr := []byte(str)
	// arr[0] = 'z'
	// str = string(arr)
	fmt.Println(str)

}