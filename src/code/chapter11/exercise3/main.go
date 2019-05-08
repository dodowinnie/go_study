package main

import (
	"fmt"

)

type Box struct {
	Len int
	Width int
	Height int
}

func (bpt *Box) getTiji() int{
	fmt.Println("请输入长方体的长")
	fmt.Scanln(&(bpt.Len))

	fmt.Println("请输入长方体的宽")
	fmt.Scanln(&(bpt.Width))

	fmt.Println("请输入长方体的高")
	fmt.Scanln(&(bpt.Height))

	return bpt.Len * bpt.Width * bpt.Height
}

func main() {
	var box Box
	tiji := box.getTiji()
	fmt.Println("长方体的体积是：", tiji)
	
}