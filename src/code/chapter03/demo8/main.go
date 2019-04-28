package main
// import "fmt"
// import "unsafe"

import(
	"fmt"
	"unsafe"
	
)

func main() {
	var b = false
	fmt.Println(b)
	fmt.Printf("b的大小=>%d", unsafe.Sizeof(b))
}