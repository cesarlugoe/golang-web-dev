package main

import (
	"io"
	"net/http"
)

// this simplifies the custom handlers and mux, just using the default
// http.handlefunc that takes a func with (res http.ResponseWriter, req *http.Request)
// NOT a type handler

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	//alternative way :
	// http.Handle("/dog", http.HandlerFunc(d))
	// http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
