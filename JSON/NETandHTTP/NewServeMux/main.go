package main

import (
	"io"
	"net/http"
)

// run main.go and go to "localhost:8080/before" and "localhost:8080/after" to see result
type bow string

type sensei string

func (s sensei) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>BOW TO YOUR SENSEI!!!</h1>")
}

func (b bow) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>OOOSSS!!!</h1>")
}

func main() {
	var s sensei
	var b bow

	mux := http.NewServeMux()
	mux.Handle("/before", s)
	mux.Handle("/after", b)
	http.ListenAndServe(":8080", mux)
}
