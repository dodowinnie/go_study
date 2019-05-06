package main

import(
	"fmt"
	"errors"
)

func test(){
	defer func(){
		err := recover() // 内置函数，可以捕获异常
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func readConf(name string) (err error) {
	if name == "config.ini"{
		return nil
	}else {
		return errors.New("读取文件错误")
	}
}

func test02(){
	err := readConf("config2.ini")
	if err != nil {
		panic(err)
	}

	fmt.Println("test02.....")
}


func main(){

	// test()
	test02()
	fmt.Println("main....")
}