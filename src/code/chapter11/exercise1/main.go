package main

import (
	"fmt"

)

type MethodUtils struct{
	
}

func (mpt *MethodUtils) JudgeNum(num int){
	if num % 2 == 0{
		fmt.Println("是偶数")
	}else {
		fmt.Println("是奇数")
	}
}

func main() {
	var m MethodUtils
	m.JudgeNum(52)
}