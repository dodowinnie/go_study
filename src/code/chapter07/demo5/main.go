package main

import (
	"fmt"
)

func swap(n1 *int, n2 *int){
	var tmp int
	tmp = *n1
	*n1 = *n2
	*n2 = tmp
}

func main(){
	n1 := 1
	n2 := 2
	swap(&n1, &n2)
	fmt.Printf("n1=%v, n2=%v", n1, n2)
}