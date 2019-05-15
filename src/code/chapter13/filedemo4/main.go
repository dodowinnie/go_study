package main

import(
	"fmt"
	// "bufio"
	"os"
	"io/ioutil"
)

func main(){
	// 复制文件
	source := "D:/opt/test/test.txt"
	dest := "D:/opt/test/dest.txt"

	strArr, err := ioutil.ReadFile(source)
	if err != nil {
		fmt.Println("read file err ", err)
		return
	}

	writerErr := ioutil.WriteFile(dest, strArr, os.ModeAppend)
	if writerErr != nil {
		fmt.Println("write file err ", writerErr)
	}



}