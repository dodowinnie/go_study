package main

import (
	"fmt"
	// "strconv"
	"strings"
)

func main (){
	// str := "hello被"
	// fmt.Println("str length is ", len(str))

	// str2 := []rune(str)
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("字符=%c\n", str2[i])
	// }

	// str to int
	// str := "123a"
	// num, err := strconv.Atoi(str)
	// if err != nil {
	// 	fmt.Println("转换错误", err)
	// }else {
	// 	fmt.Println("result is ", num)
	// }

	// int to str
	// str := strconv.Itoa(123a)
	// fmt.Printf("str type is %T, str is %v", str, str)

	// 字符串转byte
	// var bytes = []byte("hello go")
	// fmt.Printf("bytes is %v", bytes)

	// byte to str
	// str := string([]byte{97, 98, 99})
	// fmt.Printf("str is %v", str)

	// 查找子串
	// b := strings.Contains("seafood", "pp")
	// fmt.Printf("b=%t", b)

	// 统计指定字符串
	// num := strings.Count("ceheese", "e")
	// fmt.Println(num)

	// 判断字符串是否相同
	// == 区分大小写
	// strings.EqualFold() 不区分大小写
	// b := strings.EqualFold("abc", "ABC")
	// fmt.Printf("b=%t", b)

	// 返回子串在字符串第一次出现的index值，若没有返回-1
	// index := strings.Index("NTL_ABC", "ere")
	// fmt.Println(index)

	// 返回最后一次出现的位置
	// index := strings.LastIndex("go golang", "go")
	// fmt.Println(index)

	// 替换
	// str := strings.Replace("go go hello", "go", "go语言", -1)
	// fmt.Println(str)
	
	// 拆分
	// str := strings.Split("hello,world,ok", ",")
	// fmt.Println(str)

	// 大小写转换
	// str := strings.ToLower("ABC")
	// str2 := strings.ToUpper("nvc")
	// fmt.Println(str)
	// fmt.Println(str2)

	// trim相关
	// str := strings.TrimSpace(" hhhh ")
	// str := strings.TrimLeft("! hhhh !", "!")
	// str := strings.TrimRight("! hhhh !", "!")
	// str := strings.Trim("! hhhh !", "!")
	// fmt.Println(str)

	// p判断开头结尾
	// b := strings.HasPrefix("hh-dasda", "h1h")
	// fmt.Printf("b=%t", b)
	b := strings.HasSuffix("hh-dasda-das", "as1")
	fmt.Printf("b=%t", b)



}