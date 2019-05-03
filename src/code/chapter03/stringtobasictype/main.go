package main
import (
	"fmt"
	"strconv"
)

func main(){
	var str string = "true"
	var b bool

	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%t\n", b, b)

	str1 := "12345"

	var n1 int64
	n1, _ = strconv.ParseInt(str1, 10, 64)
	n2 := int(n1)
	fmt.Printf("n1 type %T n1=%d\n", n2, n2)

	str3 := "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T and f1=%f\n", f1, f1)

	h := "hello"
	var n3 int64
	n3, _ = strconv.ParseInt(h, 10, 64)
	fmt.Printf("n2 type %T and n2=%d\n", n3, n3)

	var b1 bool
	b1, _ = strconv.ParseBool(h)
	fmt.Printf("b type %T b=%t\n", b1, b1)

}