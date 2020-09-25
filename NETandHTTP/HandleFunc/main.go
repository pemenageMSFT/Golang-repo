package main

import (
	"io"
	"net/http"
)

// 1. Run main.go
// 2. in browser go to localhost:8080/bow and localhost:8080/oss

func bow(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<h1>BOW TO YOUR SENSEI!!!</h1>")
}

func oss(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<h1>OOOSSS!!!</h1>")
}

func main() {

	http.HandleFunc("/bow", bow)
	http.HandleFunc("/oss", oss)

	http.ListenAndServe(":8080", nil)
}
