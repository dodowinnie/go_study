package main

import (
	"fmt"
	
)


type Person struct {
	Name string
	
}

func (person Person)speak(){
	fmt.Printf("%v是个好人", person.Name)
}

func (person Person) jisuan() int{
	var sum int
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	return sum
}

func (person Person) jisuan2(n int) int{
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main(){
	person := Person{"brandon"}
	person.speak()
	
	sum := person.jisuan2(100)
	fmt.Println(sum)


}