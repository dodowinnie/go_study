package main

import (
	"fmt"
)


type Rect struct {
	leftUp Point
	RightDown Point
	
}

type Rect2 struct {
	leftUp *Point
	RightDown *Point
	
}


type Point struct {
	x int
	y int
}

func main(){
	
	// r1 := Rect{Point{1,2}, Point{3,4}}
	// fmt.Printf("r1.LeftUp-x=%v\n", &r1.leftUp.x)
	// fmt.Printf("r1.LeftUp-y=%v\n", &r1.leftUp.y)
	// fmt.Printf("r1.RightDown-x=%v\n", &r1.RightDown.x)
	// fmt.Printf("r1.RightDown-y=%v\n", &r1.RightDown.y)

	r2 := Rect2{&Point{10,20}, &Point{30,40}}
	fmt.Printf("r2.LeftUp=%p\n", r2.leftUp)
	fmt.Printf("r2.RightDown=%p\n", r2.RightDown)
}