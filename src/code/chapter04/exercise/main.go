package main
import (
	"fmt"
)

func main(){
	var a int = 10
	var b int = 20

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a=%v, b=%v", a, b)
}