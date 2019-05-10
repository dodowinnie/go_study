package main

import (
	"fmt"
	"os"
)

func main()  {
	file, err := os.Open("D:/opt/test/test.txt")
	if err != nil {
		fmt.Println("open err=", err)
	}

	fmt.Printf("%v", *file)

	err = file.Close()
	if err != nil {
		fmt.Println("close err=", err)
	}


}

