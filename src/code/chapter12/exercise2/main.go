package main

import (
	"fmt"
	
)

type Monster struct {
	Name string
	Age int
}

type E struct {
	Monster
	int
}


func main()  {
	e := E {
		Monster{"brandon", 28},
		100,
	}

	fmt.Println(e.int)
}