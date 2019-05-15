package main

import(
	"fmt"
	"bufio"
	"os"
	"io"
)

type CharCount struct{
	EnCount int // 英文个数
	NumCount int // 数字个数
	SpaceCount int //空格个数
	OtherCount int // 其他字符个数
}

func main(){
	fileName := "D:/opt/test/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	// 创建CharCount实例
	var count CharCount
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		strRune := ([]rune)(str)
		for _, val := range strRune {
			switch {
			case val >= 'a' && val <= 'z':
				fallthrough // 穿透
			case val >= 'A' && val <= 'Z':
				count.EnCount++
			case val == ' ' || val == '\t':
				count.SpaceCount++
			case val >= '0' && val <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}

	}

	fmt.Printf("字符的个数为=%v\n数字的个数为=%v\n空格的个数为=%v\n其他字符个数为=%v", count.EnCount, count.NumCount, count.SpaceCount, count.OtherCount)





}