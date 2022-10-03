package main

import (
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Login         string
	Password      string
	Success       bool
	StorageAccess string
}

var (
	tmpl = template.Must(template.ParseFiles("forms.html"))
)

func handler(w http.ResponseWriter, req *http.Request) {
	data := ContactDetails{
		Login:    req.FormValue("login"),
		Password: req.FormValue("Password"),
	}
	data.Success = false //true
	data.StorageAccess = "Hi User!"
	err := tmpl.Execute(w, data)

	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
