package main

import (
	"fmt"
)


type Person struct {
	Name string
	Age int
	Scores [5]int
	ptr *int
	slice []int
	map1 map[string]string
}

func main(){
	var tom Person
	fmt.Println(tom)

	// slice := []int{1}
	tom.slice = make([]int, 10, 15)
	

	tom.map1 = make(map[string]string)

	fmt.Println(tom)

	
}