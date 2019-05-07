package main

import (
	"fmt"
)


type Person struct {
	Name string
	Age int
	
}

func main(){
	// p1 := new(Person)
	// p1.Name = "brandon"
	// fmt.Println(*p1)
	p1 := &Person{}
	p1.Name = "dodo"
	fmt.Println(*p1)
	
}