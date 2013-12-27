package main

import (
	"html/template"
	"net/http"
)

var baseTemplate = "templates/_base.html"

var indexTemplate = template.Must(template.ParseFiles(
	baseTemplate,
	"templates/index.html",
))

var secondPageTemplate = template.Must(template.ParseFiles(
	baseTemplate,
	"templates/secondPageTemplate.html",
))

var thirdPageTemplate = template.Must(template.ParseFiles(
	baseTemplate,
	"templates/thirdPageTemplate.html",
))

var fourthPageTemplate = template.Must(template.ParseFiles(
	baseTemplate,
	"templates/fourthPageTemplate.html",
))

var fifthPageTemplate = template.Must(template.ParseFiles(
	baseTemplate,
	"templates/fifthPageTemplate.html",
))

func indexFunction(w http.ResponseWriter, req *http.Request) {
	if err := indexTemplate.Execute(w, nil); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func secondPageFunction(w http.ResponseWriter, req *http.Request) {
	if err := secondPageTemplate.Execute(w, nil); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func thirdPageFunction(w http.ResponseWriter, req *http.Request) {
	if err := thirdPageTemplate.Execute(w, nil); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fourthPageFunction(w http.ResponseWriter, req *http.Request) {
	if err := fourthPageTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fifthPageFunction(w http.ResponseWriter, req *http.Request) {
	if err := fifthPageTemplate.Execute(w, nil); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("templates/stylesheets")))) 

	http.HandleFunc("/index.html", indexFunction)
	http.HandleFunc("/secondPageTemplate.html", secondPageFunction)
	http.HandleFunc("/thirdPageTemplate.html", thirdPageFunction)
	http.HandleFunc("/fourthPageTemplate.html", fourthPageFunction)
	http.HandleFunc("/fifthPageTemplate.html", fifthPageFunction)

	http.HandleFunc("/", indexFunction)

	if err := http.ListenAndServe(":8080", nil); err != nil {
    	panic(err)
	}
}

