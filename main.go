package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Hello from Snippetbox")
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Display a specific snippet...")
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new snippet...")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetView)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
