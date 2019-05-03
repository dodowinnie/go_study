package main
import (
	"fmt"
)

func main () {
	i := 1
	for i; i <= 10; i++ {
		fmt.Println("hello world", i)
	}
	fmt.Println("i=", i)
}