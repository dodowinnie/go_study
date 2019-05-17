package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const tpl = `
<html>
	<head>
		<title>Hey</title>
	</head>
	<body>
		<form method="POST" action="/>
			username: <input name="username" type="text"/>
			password: <input name="password" type="password"/>
			<button type="submit">提交</button>
		</form>
	</body>
</html>

`

func main() {
	http.HandleFunc("/", Hey)
	http.ListenAndServe(":8080", nil)
}

func Hey(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.New("hey")
		t.Parse(tpl)
		t.Execute(w, nil)
	} else {
		fmt.Printf(r.FormValue("username"))
	}
}
