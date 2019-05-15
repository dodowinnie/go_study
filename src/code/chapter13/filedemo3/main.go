package main

import(
	"fmt"
	"bufio"
	"os"
	"io"
)

func main(){
	filePath := "D:/opt/test/test.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND , 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	// 及时关闭file句柄
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Println(content)
	}


	str := "hello brandon\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5;i++{
		writer.WriteString(str)
	}
	// writer.WriteString(str)
	writer.Flush()

}