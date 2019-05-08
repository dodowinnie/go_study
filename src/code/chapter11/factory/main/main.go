package main

import(
	"code/chapter11/factory/model"
	"fmt"
)

func main(){
	// var stu = model.Student{
	// 	Name:"brandon",
	// 	Score:88.6,
	// }

	// fmt.Println(stu)
	stu := model.NewStudent("brandon", 78.5)
	
	fmt.Println(stu.GetScore())
}