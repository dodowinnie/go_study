package main

import (
	"fmt"
	
)

type Integer int

func (i Integer) print() {
	fmt.Println("i=",i)
}

func (i *Integer) upadte(){
	*i = *i + 1
}


type Student struct{

	Name string
	Age string
}

func (stu *Student) String() string{
	str := fmt.Sprintf("name=[%v], age=[%v]", stu.Name, stu.Age)
	return str
}


func main(){
	stu1 := Student{"brandon", "25"}
	fmt.Println(&stu1)


}