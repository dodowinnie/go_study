package main
import (
	"fmt"
	"strconv"
)

func main(){
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string

	// fmt.Sprintf
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str = %q\n", str, str)

	// 使用strconv函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str = %q\n", str, str)

	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("str type iss %T str = %q", str, str)

}