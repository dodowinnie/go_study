package main

import (
	"fmt"
)

type A struct {
	Name string
	age int
}


type B struct{
	Name string
	score float64
}

type C struct{
	A
	B
	Name string
}


type Goods struct{
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}


type D struct{
	a A
}

func main()  {
	tv := TV{
		Goods{"电视机", 5000},
		Brand{"海尔", "山东"},
	}

	tv2 := TV{
		Goods{
			Price:1000,
			Name:"lll",
		},
		Brand{
			Name:"nnn",
			Address:"amsmd",
		},
	}

	tv3 := TV2{
		&Goods{
			Price:11000,
			Name:"内存坚实的",
		},
		&Brand{
			Name:"二级分",
			Address:"打开圣诞节",
		},
	}

	fmt.Println(tv)
	fmt.Println(tv2)
	fmt.Println(*tv3.Goods)


}