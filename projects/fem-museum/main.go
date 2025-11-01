package main

import (
	"fmt"
	"html/template"
	"net/http"

	"frontendmasters.com/go/museum/api"
	"frontendmasters.com/go/museum/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from a Go program !"))

}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	html.Execute(w, data.GetAll())
}

func main() {

	fs := http.FileServer(http.Dir("./public"))

	server := http.NewServeMux()

	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)
	server.Handle("/", fs)

	err := http.ListenAndServe(":3339", server)

	if err != nil {
		fmt.Printf("Error while opening the server %v /n", err)
	}
}
