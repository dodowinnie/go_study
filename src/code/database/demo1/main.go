package main

import (
	"io"
	"net/http"
)

func Cookie(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	http.SetCookie(w, ck)

	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}

func Cookie2(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello world",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	w.Header().Set("Set-Cookie", ck.String())

	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}

func main() {
	http.HandleFunc("/", Cookie2)

	http.ListenAndServe(":8080", nil)
}
