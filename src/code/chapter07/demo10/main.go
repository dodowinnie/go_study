package main
import (
	"fmt"
	// "encoding/json"
)

var age int = 100
var Name string = "brandon"


func test() {
	age := 10
	name := "tom"
	fmt.Println("test age=", age, "name=", name)
}

func main(){
	
	fmt.Println("main age=", age, "name=", Name)
}