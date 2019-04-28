package main
// import "fmt"
// import "unsafe"

import(
	"fmt"
	
	
)

func main() {

	var i int32 = 100
	var n1 float32 = float32(i)
	fmt.Printf("i=%v, n1=%v\n", i, n1)

	//被转换的变量存储的数据（即值），变量本身的数据类型并没有变化
	fmt.Printf("i type is %T\n", i)

	// 在转换中，比如讲int64转成int8，编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println(num2)

}