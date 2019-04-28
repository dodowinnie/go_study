package main
// import "fmt"
// import "unsafe"

import(
	"fmt"
	
)

func main() {
	var price float32 = 89.32
	fmt.Println("price is ", price)

	// 精度损失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)
	var num5 = 1.1
	fmt.Printf("num5的数据类型是%T\n", num5)

	num6 := 5.12
	num7 := .123
	fmt.Println("num6=", num6, "num7=", num7)

	var n1 = 'a' + 10
	fmt.Println("n1=", n1)

}