package main

import (
	"fmt"
	"code/chapter12/exercise/model"
)

func main()  {
	account := model.NewAccount("brandon", "123345", 60)
	fmt.Println(account.GetBalance("123345"))
}