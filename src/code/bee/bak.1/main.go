package main

import(
	// "github.com/astaxie/beego"
	"io"
	"log"
	"net/http"
)

type myHandler struct{}

func main()  {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", mux)
	if err != nil{
		log.Fatal(err)
	}
}



func (this *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "URL:" + r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello world, this is version 2.")
}