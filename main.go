package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)


func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox!"))
}


func showSnippet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	id, err := strconv.Atoi(query)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with id %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("create a snippet!"))
}


func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :2021")
	err := http.ListenAndServe(":2021", mux)
	if err != nil {
		log.Fatal("Could not start server on port 2021", err)
	}
}
