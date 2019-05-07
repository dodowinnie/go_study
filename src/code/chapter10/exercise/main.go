package main

import (
	"fmt"
)

func modifyUser(users map[string]map[string]string, username string){
	if users[username] != nil {
		users[username]["pwd"] = "888888"
	}else {
		users[username] = make(map[string]string)
		users[username]["pwd"] = "888888"
		users[username]["nickname"] = "nickname==" + username
	}
}

func main(){
	// var username string
	// var password string

	users := make(map[string]map[string]string)
	users["smith"] = make(map[string]string)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "nickname==smith"



	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "smith")

	fmt.Println(users)


}