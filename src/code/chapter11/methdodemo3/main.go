package main

import (
	"fmt"
	
)

type Person struct{
	Name string
}

func test01(p Person){
	fmt.Println(p.Name)
}

func (p Person) test02(){
	p.Name = "jack"
	fmt.Println("test02====", p.Name)
}

func (p *Person) test03(){
	p.Name = "dodo"
	fmt.Println("test02====", p.Name)
}


func main(){
	p := Person{"tom"}
	// p.test02()
	// (&p).test02()
	// fmt.Println("p.name=", p.Name)
	p.test03()
	(&p).test03()
	fmt.Println("p.name=", p.Name)


}