package main
import (
	"fmt"
)

func main () {
	// i := 1
	// for i; i <= 10; i++ {
	// 	fmt.Println("hello world", i)
	// }
	// fmt.Println("i=", i)

	// j := 1
	// for j <= 10 {
	// 	fmt.Println("hello world", j)
	// 	j++
	// }

	// m := 1
	// for{
	// 	fmt.Println("hello world", m)
	// 	m++
	// 	if m == 10 {
	// 		break
	// 	}
	// }
	// var str string = "hello world哈哈哈"
	// str2 := []rune(str)
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("%c\n", str2[i])
	// }
	str := "abc-ok哈哈"
	for index, val := range str{
		fmt.Printf("index=%d, val=%c \n", index, val)
	}



}