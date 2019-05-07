package main

import(
	"fmt"
)

func main (){
	// monster := []map[string]string
	// monster = make([]map[string]string)
	var monster []map[string]string = make([]map[string]string, 2)
	if monster[0] == nil {
		monster[0] = make(map[string]string, 2)
		monster[0]["name"] = "brandon"
		monster[0]["age"] = "500"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string, 2)
		monster[1]["name"] = "dodo"
		monster[1]["age"] = "100"
	}

	monsterNew := map[string]string{
		"name":"yy",
		"age":"200",
	}

	monster  = append(monster, monsterNew)

	fmt.Println(monster)

}