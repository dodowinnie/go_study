package main
// import "fmt"
// import "unsafe"

import(
	"fmt"
	"unsafe"
)

func main() {
	// var j int = 458
	// fmt.Println(j)

	// int, byte, unit, rune
	var a int = 8900
	// fmt.Printf("")用于做格式化输出
	fmt.Printf("a 的类型 %T", a)

	// 查询某个变量的字节大小和数据类型
	var n2 int64 = 10
	fmt.Printf("n2的数据类型 %T，n2占用的字节数是%d\n", n2, unsafe.Sizeof(n2))


	var age byte = 12
	fmt.Println("brandon的年龄是", age)





}