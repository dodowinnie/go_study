package main

import(
	"fmt"
)

func main(){
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}

	if age > 30 || age < 40{
		fmt.Println("ok2")
	}

}