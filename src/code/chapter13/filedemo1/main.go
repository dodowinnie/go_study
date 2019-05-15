package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main()  {
	file, err := os.Open("D:/opt/test/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	defer file.Close()

	// const (
	// 	default BufSize = 4096 // 默认的缓冲区4096
	// )

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF{ // io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)

	}

	fmt.Println("文件读取结束")

	// fmt.Printf("%v", *file)

	// err = file.Close()
	// if err != nil {
	// 	fmt.Println("close err=", err)
	// }


}

