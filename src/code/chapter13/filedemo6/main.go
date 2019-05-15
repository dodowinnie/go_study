package main

import(
	"fmt"
	"bufio"
	"os"
	"io"
)

func CopyFile(destFile string, sourceFile string)(written int64, err error){
	// 打开源文件
	source, err := os.Open(sourceFile)
	if err != nil {
		fmt.Printf("open file err %v", err)
		return
	}
	defer source.Close()
	// 获取reader
	reader := bufio.NewReader(source)

	// 打开目标文件
	dest, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE,0666)
	if err != nil{
		fmt.Printf("open file err %v", err)
		return
	}
	defer dest.Close()
	// 获取writer
	wirter := bufio.NewWriter(dest)
	return io.Copy(wirter, reader)
}

func main(){
	dest := "D:/opt/test/2.jpg"
	source := "D:/opt/image/20181114/5/1.jpg"
	_, err := CopyFile(dest, source)
	if err == nil{
		fmt.Println("拷贝完成")
	}else {
		fmt.Printf("拷贝失败，err:%v", err)
	}





}