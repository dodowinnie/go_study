package main

import (
	"fmt"

)

type Visitor struct {
	Age int
}

func (vpt *Visitor) getTicket() int {
	if vpt.Age > 18 {
		return 20
	}else {
		return 0
	}
}

func main() {
	var visitor Visitor
	fmt.Println("请输入年龄")
	fmt.Scanln(&visitor.Age)
	ticket := visitor.getTicket()
	fmt.Println("票价为：", ticket)

	
}