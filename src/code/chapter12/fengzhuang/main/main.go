package main
import (
	"fmt"
	"code/chapter12/fengzhuang/model"
)

func main()  {
	p := model.NewPerson("brandon")
	p.SetAge(28)
	p.SetSal(10000.15)
	fmt.Println(p.GetSal())
}