package main

import (
	"fmt"
	
)

type Circle struct {
	Radius float64
}


func (cpt *Circle) area() float64{
	fmt.Printf("%p",cpt)
	cpt.Radius = 100
	return (*cpt).Radius * (*cpt).Radius * 3.14
}


func main(){
	circle := Circle{10}
	fmt.Printf("%p",&circle)
	area := (&circle).area()
	fmt.Println(area)
	fmt.Println(circle.Radius)


}