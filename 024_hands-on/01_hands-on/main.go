package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type data struct {
	message string
	image   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `foo ran`)
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", "This is from dog")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}

}
