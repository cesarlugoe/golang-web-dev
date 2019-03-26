package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	// dog handler, handles path="/dog/something/else"
	mux.Handle("/dog/", d)
	// cat handler, doesn't take path="/cat/something/else" goes 404
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
