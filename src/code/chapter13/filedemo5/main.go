package main

import(
	"fmt"
	// "bufio"
	"os"
	// "io/ioutil"
)

func main(){
	// 复制文件
	source := "D:/opt/test/test3.txt"
	_, err := os.Stat(source)
	if err == nil {
		fmt.Println("file is exist")
	}else {
		if os.IsNotExist(err){
			fmt.Println("file is not exist")
		}
	}





}