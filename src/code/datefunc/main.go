package main

import (
	"fmt"
	"time"
	"strconv"
)

func test(){
	str := "hello"
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
}

func main(){
	// 当前时间
	begin := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("所用时间%v秒", (end - begin))

	// fmt.Printf("now=%v, now type=%T\n", now, now)
	// 年
	// fmt.Printf("year=%v,month=%v, day=%v", now.Year(), int(now.Month()), now.Weekday())

	// 格式化日期时间
	// fmt.Printf("当前时间 %d-%d-%d %d:%d:%d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	// 每隔0.5秒打印一句话
	// i := 0
	// for {
	// 	i ++
	// 	fmt.Println(i)
	// 	time.Sleep(500 * time.Millisecond)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// 获取当前unix时间戳
	// fmt.Println(now.UnixNano())




}