package main

import (
	"io"
	"net/http"
)

func i(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "index")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

func m(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cesar")
}

func main() {
	http.HandleFunc("/", i)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8080", nil)
}
