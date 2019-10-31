package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title     string
	FirstName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/register", register)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, r *http.Request) {

	pd := pageData{
		Title: "Index page",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
func about(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title: "About page",
	}
	err := tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		log.Println(err)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title: "Contact page",
	}
	err := tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		log.Println(err)
	}
}
func register(w http.ResponseWriter, r *http.Request) {
	var firstName string
	pd := pageData{
		Title: "Register page",
	}
	if r.Method == http.MethodPost {
		firstName = r.FormValue("fname")
		pd.FirstName = firstName
	}
	err := tpl.ExecuteTemplate(w, "register.gohtml", pd)
	if err != nil {
		log.Println(err)
	}
}
