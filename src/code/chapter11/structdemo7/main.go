package main

import (
	"fmt"
	
)

type Student struct{
	Name string
	Age int
}

func main(){

	stu1 := Student{"brandon", 18}
	fmt.Println(stu1)

	stu2 := Student{
		Name:"dodo",
		Age:28,
	}

	fmt.Println(stu2)

	var stu3 = Student{
		Name:"cy",
		Age:23,
	}
	
	fmt.Println(stu3)

	stu4 := &Student{
		Name :"lll",
		Age:44,
	}

	fmt.Println(*stu4)

	var stu5 = &Student{
		Name:"cy123",
		Age:43,
	}

	fmt.Println(*stu5)
}