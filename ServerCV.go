package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// ----------------------------------------------------------------------------------------
// Definition des structures de la page et de la variable global
type PageCv struct {
	Name  string
	Intro string
	Links []Link
}

type Link struct {
	Name string
	Url  string
}
var page PageCv


// ----------------------------------------------------------------------------------------
// Fonction de generation de la page
func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	data := page
	tmpl.Execute(w, data)
}

// ----------------------------------------------------------------------------------------
// Fonction Serveur
func server(port string){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// ----------------------------------------------------------------------------------------
// Main
func main() {
	// Verification du nombre de parametres
	if len(os.Args) < 2 {
		log.Fatal("Usage: PROG [Port] [Name] {[Link name 1] [link url 1]} {...} ...")
		return
	}
	// Defininition du nom
	page.Name = os.Args[2]

	// Definition des liens
	var i int = 3;
	for i < len(os.Args)-1{
		page.Links = append(page.Links, Link{Name: os.Args[i], Url: os.Args[i+1]})
		i += 2
	}
	server(os.Args[1]) // Demarage du serveur avec les arguments
}
