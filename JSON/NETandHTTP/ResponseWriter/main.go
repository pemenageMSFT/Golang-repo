package main

import (
	"fmt"
	"net/http"
)

type Sensei struct {
	Firstname string
	Lastname  string
}

var Profess = Sensei{
	Firstname: "PROFESSOR",
	Lastname:  "SENSEI",
}

func (m Sensei) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Professor-Key", "This is from Sensei")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>BOW TO YOUR: %v %v!!!</h1>", Profess.Firstname, Profess.Lastname)
}

func main() {
	http.ListenAndServe(":8080", Profess)
}
