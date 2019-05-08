package main

import (
	"fmt"
)


type Person struct {
	Name string
	Age int
	
}

func main(){
	var p1 Person
	p1.Name = "brandon"
	p1.Age = 10

	var p2 *Person = &p1

	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)
	p2.Name = "dodo"
	fmt.Printf("p2.Name=%v, p1.Name=%v\n", p2.Name, p1.Name)
	fmt.Printf("p2.Name=%v, p1.Name=%v\n", (*p2).Name, p1.Name)
	fmt.Printf("%p\n", &p1)
	fmt.Printf("%p", &p2)
	
}